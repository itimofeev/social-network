package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const idKey = "_id"

func (r *Repository) SendMessage(ctx context.Context, fromUser, toUser uuid.UUID, messageText string, ts time.Time) error {
	// todo create index on from to fields
	// todo create index on ts
	// https://stackoverflow.com/questions/16358857/mongodb-atomic-findorcreate-findone-insert-if-nonexistent-but-do-not-update
	// https://stackoverflow.com/questions/22454737/is-there-an-option-to-find-or-insert-in-mongodb

	dialogId, err := r.findDialogID(ctx, fromUser, toUser)
	if err != nil {
		return err
	}

	// todo make concurrent safe
	_, err = r.db.Collection("messages").InsertOne(ctx, messageDTO{
		ID:        primitive.NewObjectID(),
		DialogID:  dialogId.(primitive.ObjectID),
		Author:    fromUser,
		Recipient: toUser,
		Text:      messageText,
		Ts:        ts,
	})
	if err != nil {
		return fmt.Errorf("failed to insert message: %w", err)
	}

	return nil
}

func (r *Repository) findDialogID(ctx context.Context, fromUser uuid.UUID, toUser uuid.UUID) (interface{}, error) {
	user1 := fromUser
	user2 := toUser

	if user1.String() > user2.String() {
		user1, user2 = user2, user1
	}

	var result bson.M

	filter := bson.D{{Key: "user1", Value: user1}, {Key: "user2", Value: user2}}
	err := r.db.Collection("dialogs").FindOne(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		insertOneResult, _err := r.db.Collection("dialogs").InsertOne(ctx, bson.D{
			{Key: "user1", Value: user1},
			{Key: "user2", Value: user2},
		})
		if _err != nil {
			return nil, fmt.Errorf("failed to insert dialog: %w", err)
		}
		return insertOneResult.InsertedID, nil
	}

	if err != nil {
		return nil, err
	}

	if result[idKey] == nil {
		return nil, errors.New("dialog id not found")
	}

	return result[idKey], nil
}

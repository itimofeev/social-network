package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/itimofeev/social-network/internal/entity"
)

type messageDTOs []messageDTO

func (o messageDTOs) toEntities() []entity.Message {
	return lo.Map(o, func(m messageDTO, _ int) entity.Message {
		return entity.Message{
			ID:        m.ID.Hex(),
			DialogID:  m.DialogID.Hex(),
			Author:    m.Author,
			Recipient: m.Recipient,
			Text:      m.Text,
			Ts:        m.Ts,
		}
	})
}

type messageDTO struct {
	ID        primitive.ObjectID `bson:"_id"`
	DialogID  primitive.ObjectID `bson:"dialog_id"`
	Author    uuid.UUID          `bson:"author"`
	Recipient uuid.UUID          `bson:"recipient"`
	Text      string             `bson:"text"`
	Ts        time.Time          `bson:"ts"`
}

func (r *Repository) ListMessages(ctx context.Context, fromUser, toUser uuid.UUID, laterThan time.Time) ([]entity.Message, error) {
	dialogId, err := r.findDialogID(ctx, fromUser, toUser)
	if err != nil {
		return nil, err
	}

	// todo make concurrent safe
	filter := bson.M{
		"dialog_id": dialogId,
		"ts":        bson.M{"$gt": laterThan},
	}
	cursor, err := r.db.Collection("messages").Find(ctx,
		filter,
		options.Find().SetLimit(100),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list messages: %w", err)
	}

	var messages messageDTOs
	if err = cursor.All(ctx, &messages); err != nil {
		return nil, fmt.Errorf("failed to list messages: %w", err)
	}

	return messages.toEntities(), nil
}

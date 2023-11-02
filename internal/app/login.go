package app

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/go-faster/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/itimofeev/social-network/internal/entity"
	"github.com/itimofeev/social-network/internal/gen/api"
)

func (a App) LoginUser(ctx context.Context, userID api.OptUserId, password api.OptString) (string, error) {
	if !userID.Set {
		return "", errors.New("user is not set")
	}

	if !password.Set {
		return "", errors.New("password is not set")
	}

	user, err := a.repo.GetUserByUserID(ctx, string(userID.Value))
	if err != nil {
		return "", err
	}

	if err := a.checkPassword(user, password.Value); err != nil {
		return "", err
	}

	return a.makeTokenForUser(user), nil
}

func (a App) checkPassword(user entity.User, password string) error {
	unhexedPassword, err := hex.DecodeString(user.Password)
	if err != nil {
		return fmt.Errorf("failed to decode password: %w", err)
	}
	err = bcrypt.CompareHashAndPassword(unhexedPassword, []byte(password))
	if err != nil {
		return entity.ErrIncorrectPassword
	}
	return nil
}

func (a App) makeTokenForUser(user entity.User) string {
	token := paseto.NewToken()

	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))

	token.SetString("user-id", user.UserID)

	return token.V4Encrypt(a.secretKey, nil)
}

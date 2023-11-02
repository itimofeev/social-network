package app

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/itimofeev/social-network/internal/entity"
	"github.com/itimofeev/social-network/internal/gen/api"
)

func (a App) CreateUser(ctx context.Context, req api.OptUserRegisterPostReq) (entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Value.Password.Value), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}

	createRequest := entity.CreateUserRequest{
		UserID:     uuid.New().String(), // todo ilya ask question
		Password:   hex.EncodeToString(hashedPassword),
		FirstName:  req.Value.FirstName.Value,
		SecondName: req.Value.SecondName.Value,
		BirthDate:  time.Time(req.Value.Birthdate.Value),
		Biography:  req.Value.Biography.Value,
		Interests:  "",
		City:       req.Value.City.Value,
	}

	user, err := a.repo.InsertUser(ctx, createRequest)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrUserNotFound = errors.New("user not found")
var ErrIncorrectPassword = errors.New("incorrect password")

type User struct {
	ID uuid.UUID

	UserID   string
	Password string

	FirstName  string
	SecondName string
	BirthDate  time.Time
	Biography  string
	Interests  string
	City       string
}

type CreateUserRequest struct {
	UserID   string
	Password string

	FirstName  string
	SecondName string
	BirthDate  time.Time
	Biography  string
	Interests  string
	City       string
}

type Profile struct {
	FirstName string
	LastName  string
	Age       int64
	City      string
}

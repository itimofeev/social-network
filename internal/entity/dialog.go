package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       string
	DialogID string
	Author   uuid.UUID
	Text     string
	Ts       time.Time
}

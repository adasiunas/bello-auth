package model

import (
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;"`
	Email         string    `db:"email"`
	PasswordToken string    `db:"registration_token"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

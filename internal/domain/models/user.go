package models

import "time"

type User struct {
	ID           int64
	Login        string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

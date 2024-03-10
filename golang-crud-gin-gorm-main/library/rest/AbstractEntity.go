package rest

import (
	"time"
)

type AbstractEntity struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

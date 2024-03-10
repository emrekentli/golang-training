package rest

import (
	"time"
)

type AbstractEntity struct {
	Id        string `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

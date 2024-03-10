package model

import (
	"golang-crud-gin/library/rest"
	"time"
)

type Tags struct {
	rest.AbstractEntity
	Name string `gorm:"type:varchar(255)"`
}
type AbstractEntity struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

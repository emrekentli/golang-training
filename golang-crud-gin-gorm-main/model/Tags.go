package model

import (
	"golang-crud-gin/library/rest"
)

type Tags struct {
	rest.AbstractEntity
	Name string `gorm:"type:varchar(255)"`
}

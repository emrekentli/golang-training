package repository

import "golang-crud-gin/model"

type TagsRepository interface {
	Save(tags *model.Tags) (*model.Tags, error)
}

package repository

import "golang-crud-gin/model"

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tagsId string, tags model.Tags)
	Delete(tagsId string)
	FindById(tagsId string) (tags model.Tags, err error)
	FindAll() []model.Tags
}

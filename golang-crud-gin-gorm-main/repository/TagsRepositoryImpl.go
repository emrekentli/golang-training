package repository

import (
	"golang-crud-gin/model"
	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}
func (t *TagsRepositoryImpl) Save(tags *model.Tags) (*model.Tags, error) {
	result := t.Db.Create(tags)
	if result.Error != nil {
		// Handle database save error
		return nil, result.Error
	}

	return tags, nil
}

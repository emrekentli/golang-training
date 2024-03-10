package repository

import (
	"errors"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

func (t *TagsRepositoryImpl) Delete(tagsId string) {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

func (t *TagsRepositoryImpl) FindById(tagsId string) (model.Tags, error) {
	var tag model.Tags
	result := t.Db.First(&tag, "id = ?", tagsId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return tag, errors.New("tag is not found")
	} else if result.Error != nil {
		return tag, result.Error
	}

	return tag, nil
}

func (t *TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImpl) Update(tagId string, tags model.Tags) {
	result := t.Db.Model(&model.Tags{}).Where("id = ?", tagId).Updates(tags)
	helper.ErrorPanic(result.Error)
}

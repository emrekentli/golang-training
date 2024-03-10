package repository

import (
	"github.com/morkid/paginate"
	"golang-crud-gin/model"
	"gorm.io/gorm"
	"net/http"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func (t *TagsRepositoryImpl) Filter(request *http.Request) (paginate.Page, error) {
	var articles []model.Tags
	model := t.Db.Joins("tags").Model(&model.Tags{})

	pg := paginate.New()
	result := pg.With(model).Request(request).Response(&articles)
	return result, nil
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}
func (t *TagsRepositoryImpl) Save(tags *model.Tags) (*model.Tags, error) {
	result := t.Db.Create(tags)
	if result.Error != nil {
		return nil, result.Error
	}

	return tags, nil
}

func (t *TagsRepositoryImpl) Update(tags *model.Tags) (*model.Tags, error) {
	result := t.Db.Save(&tags)
	if result.Error != nil {
		return nil, result.Error
	}

	return tags, nil
}

func (t *TagsRepositoryImpl) FindById(id string) (*model.Tags, error) {
	var tags model.Tags
	tags.Id = id
	result := t.Db.First(&tags)
	if result.Error != nil {
		return nil, result.Error
	}

	return &tags, nil
}

func (t *TagsRepositoryImpl) FindAll() ([]model.Tags, error) {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	if result.Error != nil {
		return nil, result.Error
	}

	return tags, nil
}

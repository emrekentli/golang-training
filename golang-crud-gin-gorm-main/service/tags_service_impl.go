package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}

func (t *TagsServiceImpl) Create(tags request.TagsRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)

	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t *TagsServiceImpl) Delete(tagsId string) {
	t.TagsRepository.Delete(tagsId)
}

func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.ID,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

func (t *TagsServiceImpl) FindById(tagsId string) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id:   tagData.ID,
		Name: tagData.Name,
	}
	return tagResponse
}

func (t *TagsServiceImpl) Update(tagId string, tags request.TagsRequest) {
	tagData, err := t.TagsRepository.FindById(tagId)
	helper.ErrorPanic(err)

	tagData.Name = tags.Name
	t.TagsRepository.Update(tagId, tagData)
}

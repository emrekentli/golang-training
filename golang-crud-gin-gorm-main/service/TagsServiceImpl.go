package service

import (
	"golang-crud-gin/data/dto"
	"golang-crud-gin/data/request"
	MESSAGE_CODE "golang-crud-gin/library/enum"
	"golang-crud-gin/library/rest"
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

func (t *TagsServiceImpl) Create(tags *request.TagsRequest) (dto.TagsDto, rest.MetaResponse) {
	if err := t.Validate.Struct(tags); err != nil {
		return dto.TagsDto{}, rest.MetaResponseOf(MESSAGE_CODE.BAD_REQUEST.Code, err.Error())
	}

	tagModel := model.Tags{
		Name: tags.Name,
	}
	var tagDto, saveError = t.TagsRepository.Save(&tagModel)
	if saveError != nil {
		return dto.TagsDto{}, rest.MetaResponseOf(MESSAGE_CODE.BAD_REQUEST.Code, saveError.Error())
	}

	return dto.TagsDto{
		ID:        tagDto.ID,
		CreatedAt: tagDto.CreatedAt,
		UpdatedAt: tagDto.UpdatedAt,
		Name:      tagDto.Name,
	}, rest.MetaResponseSuccess()
}

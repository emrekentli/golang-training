package service

import (
	"github.com/go-playground/validator/v10"
	"golang-crud-gin/data/dto"
	"golang-crud-gin/library/constants"
	"golang-crud-gin/library/rest"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"
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

func (t *TagsServiceImpl) Create(tags *dto.TagsDto) (dto.TagsDto, rest.MetaResponse) {

	if err := t.Validate.Struct(tags); err != nil {
		var errMessage = rest.GenerateErrorMessage(&err)
		return dto.TagsDto{}, rest.MetaResponseOf(constants.BAD_REQUEST.Code, errMessage)
	}

	tagModel := model.Tags{
		Name: tags.Name,
	}

	var tagDto, saveError = t.TagsRepository.Save(&tagModel)
	if saveError != nil {
		return dto.TagsDto{}, rest.MetaResponseOf(constants.BAD_REQUEST.Code, saveError.Error())
	}

	return dto.TagsDto{
		ID:        tagDto.ID,
		CreatedAt: tagDto.CreatedAt,
		UpdatedAt: tagDto.UpdatedAt,
		Name:      tagDto.Name,
	}, rest.MetaResponseSuccess()
}

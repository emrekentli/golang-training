package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/morkid/paginate"
	"golang-crud-gin/data/dto"
	"golang-crud-gin/helper"
	"golang-crud-gin/library/constants"
	"golang-crud-gin/library/rest"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"
	"net/http"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func (t *TagsServiceImpl) Filter(request *http.Request) (paginate.Page, rest.MetaResponse) {
	var page paginate.Page
	var err error
	if err != nil {

	}
	page, err = t.TagsRepository.Filter(request)
	return page, rest.MetaResponseSuccess()
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) *TagsServiceImpl {
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

	tagModel := toEntity(&model.Tags{}, tags)

	var tagEntity, saveError = t.TagsRepository.Save(&tagModel)
	if saveError != nil {
		return dto.TagsDto{}, rest.MetaResponseOf(constants.BAD_REQUEST.Code, saveError.Error())
	}

	return toDto(tagEntity), rest.MetaResponseSuccess()
}

func (t *TagsServiceImpl) Update(id string, tags *dto.TagsDto) (dto.TagsDto, rest.MetaResponse) {
	if err := t.Validate.Struct(tags); err != nil {
		var errMessage = rest.GenerateErrorMessage(&err)
		return dto.TagsDto{}, rest.MetaResponseOf(constants.BAD_REQUEST.Code, errMessage)
	}

	currentTag, err := t.findByEntityId(id)
	if !helper.IsErrorCodeEqualsSuccess(err.Code) {
		return dto.TagsDto{}, rest.MetaResponseOf(constants.ENTITY_NOT_FOUND.Code, "Tag not found")
	}
	tagModel := toEntity(&currentTag, tags)
	var tagEntity, updateError = t.TagsRepository.Update(&tagModel)
	if updateError != nil {
		return dto.TagsDto{}, rest.MetaResponseOf(constants.BAD_REQUEST.Code, updateError.Error())
	}

	return toDto(tagEntity), rest.MetaResponseSuccess()
}

func (t *TagsServiceImpl) FindById(id string) (dto.TagsDto, rest.MetaResponse) {
	var tagEntity, updateError = t.TagsRepository.FindById(id)
	if updateError != nil {
		return dto.TagsDto{}, rest.MetaResponseOf(constants.BAD_REQUEST.Code, updateError.Error())
	}
	return toDto(tagEntity), rest.MetaResponseSuccess()
}

func (t *TagsServiceImpl) findByEntityId(id string) (model.Tags, rest.MetaResponse) {
	var tagEntity, updateError = t.TagsRepository.FindById(id)
	if updateError != nil {
		return model.Tags{}, rest.MetaResponseOf(constants.BAD_REQUEST.Code, updateError.Error())
	}
	return *tagEntity, rest.MetaResponseSuccess()
}
func (t *TagsServiceImpl) FindAll() ([]dto.TagsDto, rest.MetaResponse) {
	var tags, updateError = t.TagsRepository.FindAll()
	if updateError != nil {
		return []dto.TagsDto{}, rest.MetaResponseOf(constants.BAD_REQUEST.Code, updateError.Error())
	}
	var tagDtos []dto.TagsDto
	for _, tag := range tags {
		tagDtos = append(tagDtos, toDto(&tag))
	}
	return tagDtos, rest.MetaResponseSuccess()

}

func toEntity(entity *model.Tags, dto *dto.TagsDto) model.Tags {
	entity.Name = dto.Name
	return *entity
}

func toDto(entity *model.Tags) dto.TagsDto {
	return dto.TagsDto{
		ID:        entity.Id,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		Name:      entity.Name,
	}
}

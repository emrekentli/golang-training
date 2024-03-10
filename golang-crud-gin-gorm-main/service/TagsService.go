package service

import (
	"golang-crud-gin/data/dto"
	"golang-crud-gin/library/rest"
)

type TagsService interface {
	Create(tags *dto.TagsDto) (dto.TagsDto, rest.MetaResponse)
}

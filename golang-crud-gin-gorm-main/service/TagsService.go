package service

import (
	"golang-crud-gin/data/dto"
	"golang-crud-gin/data/request"
	"golang-crud-gin/library/rest"
)

type TagsService interface {
	Create(tags *request.TagsRequest) (dto.TagsDto, rest.MetaResponse)
}

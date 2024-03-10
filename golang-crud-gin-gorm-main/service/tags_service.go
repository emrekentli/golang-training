package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type TagsService interface {
	Create(tags request.TagsRequest)
	Update(tagId string, tags request.TagsRequest)
	Delete(tagsId string)
	FindById(tagsId string) response.TagsResponse
	FindAll() []response.TagsResponse
}

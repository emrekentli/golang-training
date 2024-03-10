package service

import (
	"github.com/morkid/paginate"
	"golang-crud-gin/data/dto"
	"golang-crud-gin/library/rest"
	"net/http"
)

type TagsService interface {
	Create(tags *dto.TagsDto) (dto.TagsDto, rest.MetaResponse)
	Update(id string, tags *dto.TagsDto) (dto.TagsDto, rest.MetaResponse)
	FindById(id string) (dto.TagsDto, rest.MetaResponse)
	FindAll() ([]dto.TagsDto, rest.MetaResponse)
	Filter(request *http.Request) (paginate.Page, rest.MetaResponse)
}

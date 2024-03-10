package repository

import (
	"github.com/morkid/paginate"
	"golang-crud-gin/model"
	"net/http"
)

type TagsRepository interface {
	Save(tags *model.Tags) (*model.Tags, error)
	Update(tags *model.Tags) (*model.Tags, error)
	FindById(id string) (*model.Tags, error)
	FindAll() ([]model.Tags, error)
	Filter(request *http.Request) (paginate.Page, error)
}

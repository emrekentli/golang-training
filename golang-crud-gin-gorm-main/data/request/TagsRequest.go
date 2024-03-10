package request

import "golang-crud-gin/data/dto"

type TagsRequest struct {
	Name string `validate:"required" json:"name"`
}

func (r *TagsRequest) ToDto() *dto.TagsDto {
	return &dto.TagsDto{
		Name: r.Name,
	}
}

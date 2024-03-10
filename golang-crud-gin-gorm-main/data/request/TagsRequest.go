package request

type TagsRequest struct {
	Name string `validate:"required" json:"name"`
}

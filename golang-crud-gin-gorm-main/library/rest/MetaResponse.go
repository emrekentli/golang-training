package rest

type MetaResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (meta *MetaResponse) newMetaResponse(code string, description string) MetaResponse {
	return MetaResponse{
		Code:        code,
		Description: description,
	}
}

func MetaResponseSuccess() MetaResponse {
	return MetaResponse{
		Code:        "200",
		Description: "SUCCESS",
	}
}
func MetaResponseOf(code string, description string) MetaResponse {
	return MetaResponse{
		Code:        code,
		Description: description,
	}
}

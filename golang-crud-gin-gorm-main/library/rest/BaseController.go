package rest

import "github.com/go-playground/validator/v10"

type BaseController struct{}

func (c *BaseController) RespondPageResponse(paginationData interface{}) PageResponseWrapper {
	return PageResponseWrapper{
		Data: paginationData,
		Meta: MetaResponseSuccess(),
	}
}

func (c *BaseController) RespondDataResponse(items interface{}) DataResponse {
	return DataResponse{
		Data: BaseDataWrapper{
			Items: items,
		},
		Meta: MetaResponseSuccess(),
	}
}

func (c *BaseController) Respond(item interface{}) ApiResponse {
	return ApiResponse{
		Data: item,
		Meta: MetaResponseSuccess(),
	}
}

func (c *BaseController) RespondError(messageCode MetaResponse) ApiResponse {
	return ApiResponse{
		Meta: MetaResponseOf(messageCode.Code, messageCode.Description),
	}
}

func GenerateErrorMessage(err *error) string {
	var errMessage string
	for _, e := range (*err).(validator.ValidationErrors) {
		errMessage = e.Error()
	}
	return errMessage
}

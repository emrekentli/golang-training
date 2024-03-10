package rest

import "github.com/go-playground/validator/v10"

type BaseController struct{}

func (c *BaseController) RespondPageResponse(items BasePageWrapper) PageResponseWrapper {
	return PageResponseWrapper{
		Data: items,
		Meta: MetaResponseSuccess(),
	}
}

func (c *BaseController) RespondDataResponse(items BaseDataWrapper) DataResponse {
	return DataResponse{
		Data: items,
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

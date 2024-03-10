package controller

import (
	"github.com/gin-gonic/gin"
	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	messagecode "golang-crud-gin/library/enum"
	"golang-crud-gin/library/rest"
	"golang-crud-gin/service"
	"net/http"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

func (controller *TagsController) Create(ctx *gin.Context) {
	createTagsRequest := request.TagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)
	var webResponse rest.ApiResponse
	var tagResponse, createError = controller.tagsService.Create(&createTagsRequest)
	if createError.Code != messagecode.SUCCESS.Code {
		webResponse = rest.ApiResponse{
			Meta: rest.MetaResponseOf(createError.Code, createError.Description),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
	} else {
		ctx.JSON(http.StatusOK, webResponse)
		webResponse = rest.ApiResponse{
			Meta: rest.MetaResponseSuccess(),
			Data: tagResponse,
		}
	}

	ctx.Header("Content-Type", "application/json")
}

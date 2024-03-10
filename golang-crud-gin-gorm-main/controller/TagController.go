package controller

import (
	"github.com/gin-gonic/gin"
	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	"golang-crud-gin/library/rest"
	"golang-crud-gin/service"
	"net/http"
)

type TagsController struct {
	rest.BaseController
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

func (controller *TagsController) Create(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	var createTagsRequest request.TagsRequest
	if err := ctx.ShouldBindJSON(&createTagsRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, controller.RespondError(rest.MetaResponseOf("400", "Invalid request")))
		return
	}
	tagResponse, createError := controller.tagsService.Create(createTagsRequest.ToDto())
	if !helper.IsErrorCodeEqualsSuccess(createError.Code) {
		ctx.JSON(http.StatusBadRequest, controller.RespondError(createError))
		return
	}
	ctx.JSON(http.StatusOK, controller.Respond(tagResponse))
}

package controller

import (
	"github.com/gin-gonic/gin"
	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	"golang-crud-gin/library/constants"
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
		ctx.JSON(http.StatusBadRequest, controller.RespondError(rest.MetaResponseOf(constants.BAD_REQUEST.Code, constants.BAD_REQUEST.Message)))
		return
	}
	tagResponse, createError := controller.tagsService.Create(createTagsRequest.ToDto())
	if !helper.IsErrorCodeEqualsSuccess(createError.Code) {
		ctx.JSON(http.StatusBadRequest, controller.RespondError(createError))
		return
	}
	ctx.JSON(http.StatusOK, controller.Respond(tagResponse))
}

func (controller *TagsController) Update(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	var updateTagsRequest request.TagsRequest
	if err := ctx.ShouldBindJSON(&updateTagsRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, controller.RespondError(rest.MetaResponseOf(constants.BAD_REQUEST.Code, constants.BAD_REQUEST.Message)))
		return
	}
	tagId := ctx.Param("id")
	tagResponse, updateError := controller.tagsService.Update(tagId, updateTagsRequest.ToDto())
	if !helper.IsErrorCodeEqualsSuccess(updateError.Code) {
		ctx.JSON(http.StatusBadRequest, controller.RespondError(updateError))
		return
	}
	ctx.JSON(http.StatusOK, controller.Respond(tagResponse))
}

func (controller *TagsController) FindById(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	tagId := ctx.Param("id")
	if tagId == "" {
		ctx.JSON(http.StatusBadRequest, controller.RespondError(rest.MetaResponseOf(constants.BAD_REQUEST.Code, constants.BAD_REQUEST.Message)))
		return
	}
	tagResponse, updateError := controller.tagsService.FindById(tagId)
	if !helper.IsErrorCodeEqualsSuccess(updateError.Code) {
		ctx.JSON(http.StatusBadRequest, controller.RespondError(updateError))
		return
	}
	ctx.JSON(http.StatusOK, controller.Respond(tagResponse))
}

func (controller *TagsController) FindAll(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	tags, findError := controller.tagsService.FindAll()
	if !helper.IsErrorCodeEqualsSuccess(findError.Code) {
		ctx.JSON(http.StatusBadRequest, controller.RespondError(findError))
		return
	}
	ctx.JSON(http.StatusOK, controller.RespondDataResponse(tags))
}

func (controller *TagsController) Filter(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	tags, findError := controller.tagsService.Filter(ctx.Request)
	if !helper.IsErrorCodeEqualsSuccess(findError.Code) {
		ctx.JSON(http.StatusBadRequest, controller.RespondError(findError))
		return
	}
	ctx.JSON(http.StatusOK, controller.RespondPageResponse(tags))
}

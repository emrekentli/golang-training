package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
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

// CreateTags		godoc
// @Summary			Create tags
// @Description		Save tags data in Db.
// @Param			tags body request.CreateTagsRequest true "Create tags"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags [post]
func (controller *TagsController) Create(ctx *gin.Context) {
	log.Info().Msg("create tags")
	createTagsRequest := request.TagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	controller.tagsService.Create(createTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateTags		godoc
// @Summary			Update tags
// @Description		Update tags data.
// @Param			tagId path string true "update tags by id"
// @Param			tags body request.CreateTagsRequest true  "Update tags"
// @Tags			tags
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagId} [patch]
func (controller *TagsController) Update(ctx *gin.Context) {
	log.Info().Msg("update tags")
	updateRequest := request.TagsRequest{}
	err := ctx.ShouldBindJSON(&updateRequest)
	helper.ErrorPanic(err)

	tagParam := ctx.Param("tagId")

	controller.tagsService.Update(tagParam, updateRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteTags		godoc
// @Summary			Delete tags
// @Description		Remove tags data by id.
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagID} [delete]
func (controller *TagsController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete tags")
	tagId := ctx.Param("tagId")
	controller.tagsService.Delete(tagId)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdTags 		godoc
// @Summary				Get Single tags by id.
// @Param				tagId path string true "update tags by id"
// @Description			Return the tahs whoes tagId valu mathes id.
// @Produce				application/json
// @Tags				tags
// @Success				200 {object} response.Response{}
// @Router				/tags/{tagId} [get]
func (controller *TagsController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid tags")
	tagId := ctx.Param("tagId")

	tagResponse := controller.tagsService.FindById(tagId)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllTags 		godoc
// @Summary			Get All tags.
// @Description		Return list of tags.
// @Tags			tags
// @Success			200 {obejct} response.Response{}
// @Router			/tags [get]
func (controller *TagsController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	tagResponse := controller.tagsService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

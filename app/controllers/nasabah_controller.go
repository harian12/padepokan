package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"padepokan/app/dto"
	"padepokan/app/services"
	"padepokan/helpers/response"
	"padepokan/helpers/validation"
	"padepokan/logger"
)

type INasabahController interface {
	GetMasabah(ctx *gin.Context)
	CreateNasabah(ctx *gin.Context)
}

type nasabahController struct {
	nasabahService services.INasabahService
}

func NewNasabahController(nasabahService services.INasabahService) *nasabahController {
	return &nasabahController{nasabahService: nasabahService}
}

func (c *nasabahController) GetMasabah(ctx *gin.Context) {

	data, err := c.nasabahService.GetNasabah()
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		ctx.JSON(http.StatusOK, response.ResponseError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.ResponseSuccess(data, "success"))
	return
}

func (c *nasabahController) CreateNasabah(ctx *gin.Context) {
	var dtNasabah dto.Nasabah

	errDTO := ctx.ShouldBindJSON(&dtNasabah)
	message, errDTO := validation.ValidationForDTO(dtNasabah)
	if errDTO != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(message))
		return
	}

	data, err := c.nasabahService.CreateNasabah(dtNasabah)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		ctx.JSON(http.StatusOK, response.ResponseError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.ResponseSuccess(data, "Data berhasil disimpan"))
	return
}

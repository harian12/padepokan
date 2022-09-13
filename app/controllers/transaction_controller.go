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

type ITransactionController interface {
	CreateTransaction(ctx *gin.Context)
}

type transactionController struct {
	transactionService services.ITransactionService
}

func NewTransactionController(transactionService services.ITransactionService) *transactionController {
	return &transactionController{transactionService: transactionService}
}

func (c *transactionController) CreateTransaction(ctx *gin.Context) {
	var dataTransaction dto.Transaction
	errDTO := ctx.ShouldBindJSON(&dataTransaction)
	message, errDTO := validation.ValidationForDTO(dataTransaction)
	if errDTO != nil {
		ctx.JSON(http.StatusOK, response.ResponseError(message))
		return
	}
	data, err := c.transactionService.CreateTransaction(dataTransaction)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		ctx.JSON(http.StatusOK, response.ResponseError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response.ResponseSuccess(data, "success"))
	return
}

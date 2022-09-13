package router

import (
	"github.com/gin-gonic/gin"
)

func RouteV1(r *gin.Engine) {
	nasabah := r.Group("nasabah")
	nasabah.GET("data", nasabah_controller.GetMasabah)
	nasabah.POST("create", nasabah_controller.CreateNasabah)

	transaction := r.Group("transaction")
	transaction.POST("create", transaction_controller.CreateTransaction)
}

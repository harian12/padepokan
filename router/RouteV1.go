package router

import (
	"github.com/gin-gonic/gin"
)

func RouteV1(r *gin.Engine) {

	v1 := r.Group("v1")
	nasabah := v1.Group("nasabah")
	nasabah.GET("data", nasabah_controller.GetMasabah)
	nasabah.GET("point", nasabah_controller.GetPointMasabah)
	nasabah.POST("create", nasabah_controller.CreateNasabah)

	transaction := v1.Group("transaction")
	transaction.POST("create", transaction_controller.CreateTransaction)
}

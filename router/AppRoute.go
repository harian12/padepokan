package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"padepokan/helpers"
)

func AppRoute() {

	appName := helpers.GetEnv("APP_NAME", "padepokan")
	appEnv := helpers.GetEnv("APP_ENV", "development")
	appPort := helpers.GetEnv("APP_PORT", "8080")

	APP := fmt.Sprintf("NAME: %s, ENV: %s", appName, appEnv)
	fmt.Println(APP)

	r := gin.Default()

	RouteV1(r)
	r.Run(":" + appPort)
}

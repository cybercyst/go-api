package main

import (
	"go-api/routes"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", routes.Ping)
	r.GET("/env", routes.Env)
	r.POST("/images", routes.UploadImage)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

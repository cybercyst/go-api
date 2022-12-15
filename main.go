package main

import (
	"go-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", routes.Ping)
	r.GET("/env", routes.Env)
	r.POST("/upload", routes.UploadFile)
	r.GET("/readyz", routes.Readyz)
	r.GET("/healthz", routes.Healthz)
	return r
}

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

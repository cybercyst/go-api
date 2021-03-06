package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Ping)
	r.GET("/env", Env)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

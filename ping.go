package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(g *gin.Context) {
	g.String(http.StatusOK, "pong")
}

package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Env(g *gin.Context) {
	g.String(http.StatusOK, os.Getenv("ENV"))
}

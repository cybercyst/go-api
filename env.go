package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Env(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"env": os.Getenv("ENV"),
		"pod": os.Getenv("PODNAME"),
	})
}

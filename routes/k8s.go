package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Readyz(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func Healthz(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"status": "OK"})
}

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(g *gin.Context) {
	g.String(http.StatusOK, "upload-image")
}

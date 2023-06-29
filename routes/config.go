package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/ffuser"
)

func Config(g *gin.Context) {
	user := ffuser.NewAnonymousUser(shortuuid.New())
	testFlag, err := ffclient.BoolVariation("test-flag", user, false)
	if err != nil {
		log.Fatal(err)
	}

	g.JSON(http.StatusOK, gin.H{
		"testFlag": testFlag,
	})
}

package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/ffcontext"
)

func Config(g *gin.Context) {
	ctx := ffcontext.NewEvaluationContextBuilder(shortuuid.New()).AddCustom("anonymous", true).Build()
	testFlag, err := ffclient.BoolVariation("test-flag", ctx, false)
	if err != nil {
		log.Fatal(err)
	}

	g.JSON(http.StatusOK, gin.H{
		"testFlag": testFlag,
	})
}

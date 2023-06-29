package main

import (
	"context"
	"go-api/routes"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/retriever/fileretriever"
)

func setupFeatureFlags() error {
	err := ffclient.Init(ffclient.Config{
		PollingInterval: 10 * time.Second,
		Logger:          log.New(os.Stdout, "", 0),
		Context:         context.Background(),
		Retriever: &fileretriever.Retriever{
			Path: "./flags.yaml",
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", routes.Ping)
	r.GET("/env", routes.Env)
	r.POST("/upload", routes.UploadFile)
	r.GET("/readyz", routes.Readyz)
	r.GET("/healthz", routes.Healthz)
	r.GET("/config", routes.Config)
	return r
}

func main() {
	err := setupFeatureFlags()
	if err != nil {
		log.Fatal(err)
	}
	defer ffclient.Close()

	r := setupRouter()
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

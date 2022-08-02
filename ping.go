package main

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func Ping(g *gin.Context) {
	done := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(time.Millisecond * 250)
	close(done)
	g.String(http.StatusOK, "pong")
}

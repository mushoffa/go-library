package http

import (
	nethttp "net/http"
	"testing"

	"github.com/mushoffa/go-library/http"

	"github.com/gin-gonic/gin"
)

func TestNewHttpServer_Success(t *testing.T) {
	handler := gin.Default()
	handler.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(nethttp.StatusOK, "TEST OK")
	})

	server := http.NewHttpServer(8084, handler)
	server.Run()
}
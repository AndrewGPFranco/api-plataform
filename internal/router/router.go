package router

import (
	"plataform/internal/handler"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handler.Ping)

	return r
}

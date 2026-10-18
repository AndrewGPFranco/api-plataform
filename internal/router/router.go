package router

import (
	"plataform/internal/handler"

	"github.com/gin-gonic/gin"
)

func New(projectHandler *handler.ProjectHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/projects", projectHandler.CreateProject)
	r.GET("/projects", projectHandler.ListProjects)
	r.DELETE("/projects/delete/:id", projectHandler.DeleteProject)
	r.PUT("/projects/edit/:id", projectHandler.EditProject)

	return r
}

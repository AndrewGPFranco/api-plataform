package router

import (
	"plataform/internal/handler"

	"github.com/gin-gonic/gin"
)

func New(projectHandler *handler.ProjectHandler, startupHandler *handler.StartupScriptHandler) *gin.Engine {
	r := gin.Default()

	// Projects
	r.POST("/projects", projectHandler.CreateProject)
	r.GET("/projects/:idUser", projectHandler.ListProjects)
	r.DELETE("/projects/delete/:id", projectHandler.DeleteProject)
	r.PUT("/projects/edit/:id", projectHandler.EditProject)

	// Scripts
	r.POST("/scripts", startupHandler.NewScript)
	r.POST("/scripts/start/:id/*path", startupHandler.StartProjectLocal)
	r.POST("/scripts/stop/:id/*path", startupHandler.EncerraProjetoLocal)

	return r
}

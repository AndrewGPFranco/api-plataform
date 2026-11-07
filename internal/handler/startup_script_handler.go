package handler

import (
	"net/http"

	"plataform/internal/dtos/scripts"
	services "plataform/internal/service/projects"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StartupScriptHandler struct {
	service *services.StartupScriptService
}

func NewStartupScriptHandler(service *services.StartupScriptService) *StartupScriptHandler {
	return &StartupScriptHandler{service: service}
}

func (h *StartupScriptHandler) NewScript(context *gin.Context) {
	var req scripts.InputStartupScript

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resposta, err := h.service.NewScript(context, req)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": resposta})
}

func (h *StartupScriptHandler) StartProjectLocal(c *gin.Context) {
	id := c.Param("id")
	path := c.Param("path")
	idParse := uuid.MustParse(id)

	response, err := h.service.StartProjectLocalhost(c, idParse, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

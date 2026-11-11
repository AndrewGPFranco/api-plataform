package handler

import (
	"errors"
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
	idParse, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id inválido"})
		return
	}

	response, err := h.service.StartProjectLocalhost(c, idParse, c.Param("path"))
	if err != nil {
		c.JSON(statusDoErroScript(err), gin.H{"message": response})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": response})
}

func (h *StartupScriptHandler) EncerraProjetoLocal(c *gin.Context) {
	idParse, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id inválido"})
		return
	}

	response, err := h.service.EncerraProcesso(idParse, c.Param("path"))
	if err != nil {
		c.JSON(statusDoErroScript(err), gin.H{"message": response})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": response})
}

func statusDoErroScript(err error) int {
	switch {
	case errors.Is(err, services.ErrScriptEmExecucao):
		return http.StatusConflict
	case errors.Is(err, services.ErrScriptSemExecucao):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

package handler

import (
	"log"
	"net/http"
	"plataform/internal/dtos"
	"plataform/internal/service"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func New(service *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req dtos.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	project := dtos.ToEntity(req)

	if err := h.service.CreateProject(c.Request.Context(), project); err != nil {
		log.Printf("erro ao criar projeto: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "erro ao registrar projeto"})
		return
	}

	c.JSON(http.StatusCreated, project)
}

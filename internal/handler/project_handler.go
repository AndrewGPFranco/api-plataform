package handler

import (
	"log"
	"net/http"
	"time"

	"plataform/internal/model"
	"plataform/internal/service"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func New(service *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

type createProjectRequest struct {
	Title          string  `json:"title" binding:"required"`
	Description    *string `json:"description"`
	StartedAt      string  `json:"started_at" binding:"required,datetime=2006-01-02"`
	RepoFrontURL   *string `json:"repo_front_url" binding:"omitempty,url"`
	RepoBackURL    *string `json:"repo_back_url" binding:"omitempty,url"`
	DeployFrontURL *string `json:"deploy_front_url" binding:"omitempty,url"`
	DeployBackURL  *string `json:"deploy_back_url" binding:"omitempty,url"`
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req createProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	startedAt, _ := time.Parse(time.DateOnly, req.StartedAt)

	project := &model.Project{
		Title:          req.Title,
		Description:    req.Description,
		StartedAt:      startedAt,
		RepoFrontURL:   req.RepoFrontURL,
		RepoBackURL:    req.RepoBackURL,
		DeployFrontURL: req.DeployFrontURL,
		DeployBackURL:  req.DeployBackURL,
	}

	if err := h.service.CreateProject(c.Request.Context(), project); err != nil {
		log.Printf("erro ao criar projeto: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "erro ao registrar projeto"})
		return
	}

	c.JSON(http.StatusCreated, project)
}

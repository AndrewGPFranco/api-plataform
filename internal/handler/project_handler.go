package handler

import (
	"log"
	"net/http"
	dtos "plataform/internal/dtos/projects"
	services "plataform/internal/service/projects"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProjectHandler struct {
	service *services.ProjectService
}

func New(service *services.ProjectService) *ProjectHandler {
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

func (h *ProjectHandler) ListProjects(c *gin.Context) {
	projects, err := h.service.ListProjects(c.Request.Context())
	if err != nil {
		log.Printf("erro ao listar projetos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "erro ao listar projetos"})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	id := c.Param("id")

	idParse := uuid.MustParse(id)

	err := h.service.DeleteProject(idParse)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Não foi possivel excluir o projeto."})
		return
	}

	c.JSON(http.StatusOK, "Projeto deletado com sucesso!")
}

func (h *ProjectHandler) EditProject(c *gin.Context) {
	var req dtos.EditProject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.service.EditProject(&req); err != nil {
		log.Printf("erro ao atualizar projeto: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Ocorreu um erro ao atualizar os metadados."})
		return
	}

	c.JSON(http.StatusOK, "Metadados atualizados.")
}

package projects

import (
	"plataform/internal/model/projects"
	"time"
)

type CreateProjectRequest struct {
	Title          string  `json:"title" binding:"required"`
	Description    *string `json:"description"`
	StartedAt      string  `json:"started_at" binding:"required,datetime=2006-01-02"`
	RepoFrontURL   *string `json:"repo_front_url" binding:"omitempty,url"`
	RepoBackURL    *string `json:"repo_back_url" binding:"omitempty,url"`
	DeployFrontURL *string `json:"deploy_front_url" binding:"omitempty,url"`
	DeployBackURL  *string `json:"deploy_back_url" binding:"omitempty,url"`
	LocalhostURL   *string `json:"localhost_url" binding:"omitempty,url"`
}

func ToEntity(dto CreateProjectRequest) *projects.Project {
	startedAt, _ := time.Parse(time.DateOnly, dto.StartedAt)

	project := &projects.Project{
		Title:          dto.Title,
		Description:    dto.Description,
		StartedAt:      startedAt,
		RepoFrontURL:   dto.RepoFrontURL,
		RepoBackURL:    dto.RepoBackURL,
		DeployFrontURL: dto.DeployFrontURL,
		DeployBackURL:  dto.DeployBackURL,
		LocalhostURL:   dto.LocalhostURL,
	}

	return project
}

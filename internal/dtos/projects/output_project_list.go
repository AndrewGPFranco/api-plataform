package projects

import (
	"time"

	"plataform/internal/model/projects"

	"github.com/google/uuid"
)

type OutputProjectList struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	Description    *string   `json:"description"`
	StartedAt      time.Time `json:"started_at"`
	RepoFrontURL   *string   `json:"repo_front_url"`
	RepoBackURL    *string   `json:"repo_back_url"`
	DeployFrontURL *string   `json:"deploy_front_url"`
	DeployBackURL  *string   `json:"deploy_back_url"`
	LocalhostURL   *string   `json:"localhost_url"`
	HasDeploy      bool      `json:"has_deploy"`
	IsOnline       bool      `json:"is_online"`
}

func FromEntity(entity projects.Project, isOnline bool) OutputProjectList {
	return OutputProjectList{
		ID:             entity.ID,
		Title:          entity.Title,
		Description:    entity.Description,
		StartedAt:      entity.StartedAt,
		RepoFrontURL:   entity.RepoFrontURL,
		RepoBackURL:    entity.RepoBackURL,
		DeployFrontURL: entity.DeployFrontURL,
		DeployBackURL:  entity.DeployBackURL,
		LocalhostURL:   entity.LocalhostURL,
		HasDeploy:      entity.HasDeploy,
		IsOnline:       isOnline,
	}
}

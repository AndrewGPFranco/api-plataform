package dtos

import (
	"plataform/internal/model"
	"time"

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
	HasDeploy      bool      `json:"has_deploy"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FromEntity(entities []model.Project) []OutputProjectList {
	list := make([]OutputProjectList, len(entities))

	for index, project := range entities {
		list[index] = OutputProjectList{
			ID:             project.ID,
			Title:          project.Title,
			Description:    project.Description,
			StartedAt:      project.StartedAt,
			RepoFrontURL:   project.RepoFrontURL,
			RepoBackURL:    project.RepoBackURL,
			DeployFrontURL: project.DeployFrontURL,
			DeployBackURL:  project.DeployBackURL,
			HasDeploy:      project.HasDeploy,
			CreatedAt:      project.CreatedAt,
			UpdatedAt:      project.UpdatedAt,
		}
	}

	return list
}

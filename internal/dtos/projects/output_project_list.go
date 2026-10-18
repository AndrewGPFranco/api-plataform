package projects

import (
	"plataform/internal/model/projects"
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
	LocalhostURL   *string   `json:"localhost_url"`
	HasDeploy      bool      `json:"has_deploy"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FromEntity(entities []projects.Project) []OutputProjectList {
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
			LocalhostURL:   project.LocalhostURL,
			HasDeploy:      project.HasDeploy,
			CreatedAt:      project.CreatedAt,
			UpdatedAt:      project.UpdatedAt,
		}
	}

	return list
}

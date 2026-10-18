package projects

import "github.com/google/uuid"

type EditProject struct {
	Id             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	Description    *string   `json:"description"`
	StartedAt      string    `json:"started_at"`
	RepoFrontURL   *string   `json:"repo_front_url" binding:"omitempty,url"`
	RepoBackURL    *string   `json:"repo_back_url" binding:"omitempty,url"`
	DeployFrontURL *string   `json:"deploy_front_url" binding:"omitempty,url"`
	DeployBackURL  *string   `json:"deploy_back_url" binding:"omitempty,url"`
	LocalhostURL   *string   `json:"localhost_url" binding:"omitempty,url"`
}

package projects

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
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

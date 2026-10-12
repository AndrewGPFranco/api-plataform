package repository

import (
	"context"
	"database/sql"

	"plataform/internal/model"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) CreateProject(ctx context.Context, project *model.Project) error {
	query := `INSERT INTO projects (title, description, started_at, repo_front_url, repo_back_url,
                      deploy_front_url, deploy_back_url)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)
			  RETURNING id, has_deploy, created_at, updated_at`

	return r.db.QueryRowContext(ctx, query,
		project.Title, project.Description, project.StartedAt, project.RepoFrontURL, project.RepoBackURL,
		project.DeployFrontURL, project.DeployBackURL,
	).Scan(&project.ID, &project.HasDeploy, &project.CreatedAt, &project.UpdatedAt)
}

package projects

import (
	"context"
	"database/sql"
	"plataform/internal/model/projects"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) CreateProject(ctx context.Context, project *projects.Project) error {
	query := `INSERT INTO projects (title, description, started_at, repo_front_url, repo_back_url,
                      deploy_front_url, deploy_back_url)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)
			  RETURNING id, has_deploy, created_at, updated_at`

	return r.db.QueryRowContext(ctx, query,
		project.Title, project.Description, project.StartedAt, project.RepoFrontURL, project.RepoBackURL,
		project.DeployFrontURL, project.DeployBackURL,
	).Scan(&project.ID, &project.HasDeploy, &project.CreatedAt, &project.UpdatedAt)
}

func (r *ProjectRepository) ListProjects(ctx context.Context) ([]projects.Project, error) {
	query := `SELECT id, title, description, started_at, repo_front_url, repo_back_url,
                     deploy_front_url, deploy_back_url, has_deploy, created_at, updated_at
			  FROM projects
			  ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	currentProjects := []projects.Project{}
	for rows.Next() {
		var p projects.Project
		if err := rows.Scan(
			&p.ID, &p.Title, &p.Description, &p.StartedAt, &p.RepoFrontURL, &p.RepoBackURL,
			&p.DeployFrontURL, &p.DeployBackURL, &p.HasDeploy, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		currentProjects = append(currentProjects, p)
	}

	return currentProjects, rows.Err()
}

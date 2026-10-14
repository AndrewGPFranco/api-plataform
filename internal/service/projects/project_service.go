package projects

import (
	"context"
	"plataform/internal/dtos/projects"
	projects2 "plataform/internal/model/projects"
	projects3 "plataform/internal/repository/projects"
)

type ProjectService struct {
	repo *projects3.ProjectRepository
}

func NewProjectService(repo *projects3.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) CreateProject(ctx context.Context, project *projects2.Project) error {
	return s.repo.CreateProject(ctx, project)
}

func (s *ProjectService) ListProjects(ctx context.Context) ([]projects.OutputProjectList, error) {
	projectsCurrents, err := s.repo.ListProjects(ctx)
	if err != nil {
		return nil, err
	}

	return projects.FromEntity(projectsCurrents), nil
}

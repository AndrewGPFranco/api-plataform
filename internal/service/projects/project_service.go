package projects

import (
	"context"
	"plataform/internal/dtos/projects"
	dtos "plataform/internal/dtos/projects"
	model "plataform/internal/model/projects"
	repository "plataform/internal/repository/projects"

	"github.com/google/uuid"
)

type ProjectService struct {
	repo *repository.ProjectRepository
}

func NewProjectService(repo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) CreateProject(ctx context.Context, project *model.Project) error {
	return s.repo.CreateProject(ctx, project)
}

func (s *ProjectService) ListProjects(ctx context.Context) ([]projects.OutputProjectList, error) {
	projectsCurrents, err := s.repo.ListProjects(ctx)
	if err != nil {
		return nil, err
	}

	return projects.FromEntity(projectsCurrents), nil
}

func (s *ProjectService) DeleteProject(id uuid.UUID) error {
	return s.repo.DeleteProject(id)
}

func (s *ProjectService) EditProject(req *dtos.EditProject) error {
	return s.repo.EditProject(req)
}

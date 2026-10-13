package service

import (
	"context"

	"plataform/internal/dtos"
	"plataform/internal/model"
	"plataform/internal/repository"
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

func (s *ProjectService) ListProjects(ctx context.Context) ([]dtos.OutputProjectList, error) {
	projects, err := s.repo.ListProjects(ctx)
	if err != nil {
		return nil, err
	}

	return dtos.FromEntity(projects), nil
}

package projects

import (
	"context"
	"fmt"
	"net/http"

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

func (s *ProjectService) ListProjects(ctx context.Context) ([]dtos.OutputProjectList, error) {
	projectsCurrents, err := s.repo.ListProjects(ctx)
	if err != nil {
		return nil, err
	}

	projects, err := handleProjects(projectsCurrents)
	if err != nil {
		fmt.Printf("Erro ao parsear o DTO para entidade %v", err)
		return projects, nil
	}

	return projects, nil
}

func handleProjects(currents []model.Project) ([]dtos.OutputProjectList, error) {
	tratado := make([]dtos.OutputProjectList, len(currents))

	for proj := range currents {
		project := currents[proj]
		_, err := http.Get(*project.LocalhostURL)

		if err == nil {
			entity := dtos.FromEntity(project, true)
			tratado[proj] = entity
		} else {
			entity := dtos.FromEntity(project, false)
			tratado[proj] = entity
		}
	}

	return tratado, nil
}

func (s *ProjectService) DeleteProject(id uuid.UUID) error {
	return s.repo.DeleteProject(id)
}

func (s *ProjectService) EditProject(req *dtos.EditProject) error {
	return s.repo.EditProject(req)
}

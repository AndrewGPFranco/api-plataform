package projects

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sync"
	"time"

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

	projects, err := handleProjects(ctx, projectsCurrents)
	if err != nil {
		fmt.Printf("Erro ao parsear o DTO para entidade %v", err)
		return projects, nil
	}

	return projects, nil
}

func handleProjects(ctx context.Context, currents []model.Project) ([]dtos.OutputProjectList, error) {
	tratado := make([]dtos.OutputProjectList, len(currents))

	client := &http.Client{Timeout: 3 * time.Second}

	var wg sync.WaitGroup

	for i := range currents {
		wg.Add(1)
		go func(i int, project model.Project) {
			defer wg.Done()

			tratado[i] = dtos.FromEntity(project, isOnline(ctx, client, project.LocalhostURL))
		}(i, currents[i])
	}

	wg.Wait()

	return tratado, nil
}

func isOnline(ctx context.Context, client *http.Client, url *string) bool {
	if url == nil {
		return false
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, *url, nil)
	if err != nil {
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		return false
	}

	_ = resp.Body.Close()

	return true
}

func (s *ProjectService) DeleteProject(id uuid.UUID) error {
	return s.repo.DeleteProject(id)
}

func (s *ProjectService) EditProject(req *dtos.EditProject) error {
	return s.repo.EditProject(req)
}

func (s *ProjectService) StartProjectLocalhost(ctx context.Context, id uuid.UUID) (string, error) {
	command, err := s.repo.FindCommandByProjectID(ctx, id)
	if err != nil {
		fmt.Println("Não foi possível encontrar o comando para a subida do projeto.")
		return "Não foi possível encontrar o comando para a subida do projeto.", err
	}

	out, err := exec.Command("sh", "-c", *command).CombinedOutput()
	if err != nil {
		log.Printf("Erro ao subir a aplicação: %v\nsaida: %s", err, out)
		return "Não foi possível iniciar o ambiente local.", err
	}
	fmt.Println(string(out))
	return "Ambiente local inicializado.", nil
}

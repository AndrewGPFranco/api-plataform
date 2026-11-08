package projects

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	dtos "plataform/internal/dtos/projects"
	model "plataform/internal/model/projects"
	repository "plataform/internal/repository/projects"

	"github.com/google/uuid"
)

type ProjectService struct {
	repo           *repository.ProjectRepository
	repoScript     *repository.StartupScriptRepository
	scriptExecucao *ScriptsEmExecucao
}

func NewProjectService(repo *repository.ProjectRepository, repoScript *repository.StartupScriptRepository, scriptExecucao *ScriptsEmExecucao) *ProjectService {
	return &ProjectService{repo: repo, repoScript: repoScript, scriptExecucao: scriptExecucao}
}

func (s *ProjectService) CreateProject(ctx context.Context, project *model.Project) error {
	return s.repo.CreateProject(ctx, project)
}

func (s *ProjectService) ListProjects(ctx context.Context) ([]dtos.OutputProjectList, error) {
	projectsCurrents, err := s.repo.ListProjects(ctx)
	if err != nil {
		return nil, err
	}

	projects, err := s.handleProjects(ctx, projectsCurrents)
	if err != nil {
		fmt.Printf("Erro ao parsear o DTO para entidade %v", err)
		return projects, nil
	}

	return projects, nil
}

func (s *ProjectService) handleProjects(ctx context.Context, currents []model.Project) ([]dtos.OutputProjectList, error) {
	tratado := make([]dtos.OutputProjectList, len(currents))

	const timeoutHTTP = 3 * time.Second
	client := &http.Client{Timeout: timeoutHTTP}

	var wg sync.WaitGroup

	for i := range currents {
		wg.Add(1)
		go func(i int, project model.Project) {
			defer wg.Done()

			scripts, _ := s.repoScript.FindScriptsByProjectID(ctx, project.ID)
			for i := range scripts {
				scripts[i].Running = s.verificaScriptJaEmExecucao(project.ID, scripts[i].Path)
			}
			tratado[i] = dtos.FromEntity(project, isOnline(ctx, client, project.LocalhostURL), scripts)
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

func (s *ProjectService) verificaScriptJaEmExecucao(idParse uuid.UUID, pathAtual *string) bool {
	if pathAtual == nil {
		return false
	}
	return s.scriptExecucao.Rodando(idParse, *pathAtual)
}

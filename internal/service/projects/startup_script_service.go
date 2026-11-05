package projects

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"plataform/internal/dtos/scripts"
	repository "plataform/internal/repository/projects"

	"github.com/google/uuid"
)

type StartupScriptService struct {
	repo *repository.StartupScriptRepository
}

func NewStartupScriptService(repo *repository.StartupScriptRepository) *StartupScriptService {
	return &StartupScriptService{repo: repo}
}

func (s *StartupScriptService) NewScript(ctx context.Context, dto scripts.InputStartupScript) (string, error) {
	err := s.repo.NewScript(ctx, dto)
	if err != nil {
		return "Ocorreu um erro ao salvar o script para o projeto.", err
	}

	return "Script salvo com sucesso", nil
}

func (s *StartupScriptService) StartProjectLocalhost(ctx context.Context, id uuid.UUID, p string) (string, error) {
	command, path, err := s.repo.FindCommandByProjectIDAndPath(ctx, id, p)
	if err != nil || command == nil {
		fmt.Println("Não foi possível encontrar o comando para a subida do projeto.")
		return "Não foi possível encontrar o comando para a subida do projeto.", err
	}

	cmd := exec.Command("sh", "-c", *command)
	if path != nil && *path != "" {
		cmd.Dir = *path
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Printf("Erro ao subir a aplicação: %v", err)
		return "Não foi possível iniciar o ambiente local.", err
	}

	pid := cmd.Process.Pid
	log.Printf("Projeto %s iniciado (pid %d) com: %s", id, pid, *command)

	go func() {
		if err := cmd.Wait(); err != nil {
			log.Printf("Projeto %s (pid %d) finalizou com erro: %v", id, pid, err)
			return
		}
		log.Printf("Projeto %s (pid %d) finalizado.", id, pid)
	}()

	return "Ambiente local inicializado.", nil
}

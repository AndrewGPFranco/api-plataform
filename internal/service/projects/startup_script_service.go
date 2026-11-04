package projects

import (
	"context"
	"fmt"
	"log"
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

func (s *StartupScriptService) StartProjectLocalhost(ctx context.Context, id uuid.UUID) (string, error) {
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

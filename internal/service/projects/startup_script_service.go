package projects

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"plataform/internal/dtos/scripts"
	repository "plataform/internal/repository/projects"

	"github.com/google/uuid"
)

var (
	ErrScriptEmExecucao  = errors.New("script já em execução")
	ErrScriptSemExecucao = errors.New("script não está em execução")
)

type StartupScriptService struct {
	repo           *repository.StartupScriptRepository
	scriptExecucao *ScriptsEmExecucao
}

func NewStartupScriptService(repo *repository.StartupScriptRepository, scriptExecucao *ScriptsEmExecucao) *StartupScriptService {
	return &StartupScriptService{repo: repo, scriptExecucao: scriptExecucao}
}

func (s *StartupScriptService) NewScript(ctx context.Context, dto scripts.InputStartupScript) (string, error) {
	err := s.repo.NewScript(ctx, dto)
	if err != nil {
		return "Ocorreu um erro ao salvar o script para o projeto.", err
	}

	return "Script salvo com sucesso", nil
}

func (s *StartupScriptService) StartProjectLocalhost(ctx context.Context, id uuid.UUID, p string) (string, error) {
	if s.scriptExecucao.Rodando(id, p) {
		return "Script já em execução", ErrScriptEmExecucao
	}

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
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	if err := cmd.Start(); err != nil {
		log.Printf("Erro ao subir a aplicação: %v", err)
		return "Não foi possível iniciar o ambiente local.", err
	}

	pid := cmd.Process.Pid

	s.scriptExecucao.Marcar(id, p, pid)

	log.Printf("Projeto %s iniciado (pid %d) com: %s", id, pid, *command)

	go func() {
		defer s.scriptExecucao.Remover(id, p)
		_ = cmd.Wait()
		log.Printf("Projeto %s (pid %d) finalizado.", id, pid)
	}()

	return "Ambiente local inicializado.", nil
}

func (s *StartupScriptService) EncerraProcesso(id uuid.UUID, p string) (string, error) {
	pid, ok := s.scriptExecucao.DevolvePidProcesso(id, p)
	if !ok {
		return "Script não está em execução", ErrScriptSemExecucao
	}

	if err := syscall.Kill(-pid, syscall.SIGTERM); err != nil {
		log.Printf("Erro ao encerrar o projeto %s (pid %d): %v", id, pid, err)
		return "Não foi possível encerrar o processo.", err
	}

	return "Processo encerrado.", nil
}

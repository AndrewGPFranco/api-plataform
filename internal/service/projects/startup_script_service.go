package projects

import repository "plataform/internal/repository/projects"

type StartupScriptService struct {
	repo *repository.StartupScriptRepository
}

func NewStartupScriptService(repo *repository.StartupScriptRepository) *StartupScriptService {
	return &StartupScriptService{repo: repo}
}

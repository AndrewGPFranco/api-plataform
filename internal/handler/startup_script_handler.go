package handler

import services "plataform/internal/service/projects"

type StartupScriptHandler struct {
	service *services.StartupScriptService
}

func NewStartupScriptHandler(service *services.StartupScriptService) *StartupScriptHandler {
	return &StartupScriptHandler{service: service}
}

package main

import (
	"log"

	proj_repos "plataform/internal/repository/projects"
	proj_services "plataform/internal/service/projects"

	"plataform/internal/config"
	"plataform/internal/database"
	"plataform/internal/handler"
	"plataform/internal/router"
)

func main() {
	cfg := config.Load()

	scriptsEmExecucao := proj_services.NewScriptsEmExecucao()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("erro ao conectar no banco: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("erro ao fechar conexão com o banco: %v", err)
		}
	}()

	projectRepo := proj_repos.NewProjectRepository(db)
	startupScriptRepo := proj_repos.NewStartupScriptRepository(db)
	startupScriptService := proj_services.NewStartupScriptService(startupScriptRepo, scriptsEmExecucao)
	startupScriptHandler := handler.NewStartupScriptHandler(startupScriptService)
	projectService := proj_services.NewProjectService(projectRepo, startupScriptRepo, scriptsEmExecucao)
	projectHandler := handler.NewProjectHandler(projectService)

	r := router.New(projectHandler, startupScriptHandler)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}

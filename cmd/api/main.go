package main

import (
	"log"

	"plataform/internal/repository/projects"
	projects2 "plataform/internal/service/projects"

	"plataform/internal/config"
	"plataform/internal/database"
	"plataform/internal/handler"
	"plataform/internal/router"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	projectRepo := projects.NewProjectRepository(db)
	projectService := projects2.NewProjectService(projectRepo)
	projectHandler := handler.New(projectService)
	startupScriptRepo := projects.NewStartupScriptRepository(db)
	startupScriptService := projects2.NewStartupScriptService(startupScriptRepo)
	startupScriptHandler := handler.NewStartupScriptHandler(startupScriptService)

	r := router.New(projectHandler, startupScriptHandler)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}

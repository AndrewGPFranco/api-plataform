package main

import (
	"log"

	"plataform/internal/config"
	"plataform/internal/database"
	"plataform/internal/handler"
	"plataform/internal/repository"
	"plataform/internal/router"
	"plataform/internal/service"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	projectRepo := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepo)
	projectHandler := handler.New(projectService)

	r := router.New(projectHandler)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}

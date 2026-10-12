package main

import (
	"log"
	"plataform/internal/config"
	"plataform/internal/database"
	"plataform/internal/router"
)

func main() {
	cfg := config.Load()
	r := router.New()

	database.CarregaBanco()

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}

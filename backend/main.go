package main

import (
	"fmt"
	"jade-grading/config"
	"jade-grading/database"
	"jade-grading/routes"
	"log"
)

func main() {
	cfg := config.Load()

	if err := database.Init(&cfg.Database); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := routes.SetupRouter()

	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on http://localhost%s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

package main

import (
	"log"

	"github.com/lipaysamart/go-todolist-api-execrices/internal/bootstrap"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/model"
	"github.com/lipaysamart/go-todolist-api-execrices/pkg/config"
	"github.com/lipaysamart/go-todolist-api-execrices/pkg/db"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.NewDatabase(cfg.DatabaseURI)
	if err != nil {
		log.Fatal("Failed serve database:", err)
	}
	if err := database.Migrate(&model.User{}, &model.Item{}); err != nil {
		log.Fatal("Failed to migrate the database:", err)
	}

	serve := bootstrap.NewBootStrap(database)
	if err := serve.Run(); err != nil {
		log.Fatal("Failed to start the server:", err)
	}

}

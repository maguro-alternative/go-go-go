package main

import (
	"fmt"
	"maguro-alternative/go-go-go/models"
	"maguro-alternative/go-go-go/pkg/db"
	"maguro-alternative/go-go-go/repository"
	"maguro-alternative/go-go-go/routes"
)

func main() {
	db, err := db.ConnectionDB()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}
	repo := repository.NewRepository(db)
	r := route.Routes(repo)
	r.Run(":8080")
}

package main

import (
	"fmt"
	"maguro-alternative/go-go-go/model"
	"maguro-alternative/go-go-go/pkg/db"
	"maguro-alternative/go-go-go/route"
)

func main() {
	db, err := db.ConnectionDB()

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Migrate the schema
	err = db.AutoMigrate(&model.Todo{})
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}
	r := route.Routes()
	r.Run(":8080")
}

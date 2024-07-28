package main

import (
	"log"
	"myblog/internal/routes"
	"myblog/pkg/db"
)

func main() {
	db.Init()
	r := routes.SetupRouter()
	log.Println("Starting server on :8080...")
	r.Run(":8080")
}

package main

import (
	"github.com/RigottiG/go-books-api/database"
	"github.com/RigottiG/go-books-api/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}

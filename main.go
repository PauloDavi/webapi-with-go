package main

import (
	"github.com/PauloDavi/webapi-with-go/database"
	"github.com/PauloDavi/webapi-with-go/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
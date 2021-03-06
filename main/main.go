package main

import (
	"fmt"
	"pizzaOrder/database"
	"pizzaOrder/server"
)

func main() {
	// Inicia o banco de dados
	database.StartDB()
	// Cria uma estância do server
	server := server.NewServer()
	// Inicia o server
	server.Run()

	fmt.Println("Server rodando na porta 5000")
}
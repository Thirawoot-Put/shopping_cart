package main

import (
	"Thirawoot/shopping_cart/internal/infrastructure/database"
	"Thirawoot/shopping_cart/internal/infrastructure/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to read env file")
	}

	db := database.Connect()

	app := server.NewServer(db)
	app.Start()
}

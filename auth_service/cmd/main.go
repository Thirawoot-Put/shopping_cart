package main

import (
	"Thirawoot/shopping_cart/internal/infrastructure/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to read env file")
	}
	database.Connect()
}

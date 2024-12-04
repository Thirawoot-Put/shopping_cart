package server

import (
	"Thirawoot/shopping_cart/internal/infrastructure/database"
	"net/http"
)

func Start() {
	database.Connect()
	HttpListen()
}

func HttpListen() {
	err := http.ListenAndServe(":9991", nil)
	if err != nil {
		panic("Failed to listen http")
	}
}

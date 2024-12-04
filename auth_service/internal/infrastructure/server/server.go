package server

import (
	"Thirawoot/shopping_cart/internal/infrastructure/database"
	"fmt"
)

func Start() {
	database.Connect()
	HttpListen()
}

func HttpListen() {
	fmt.Println("Server connect to http listen port")
}

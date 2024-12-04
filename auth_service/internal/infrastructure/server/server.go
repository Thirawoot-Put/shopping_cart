package server

import (
	"net/http"
	"os"

	"gorm.io/gorm"
)

type Server struct {
	app      *http.ServeMux
	database *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	netApp := http.NewServeMux()

	return &Server{app: netApp, database: db}
}

func (s *Server) Start() {
	port := os.Getenv("SERVER_PORT")
	s.HttpListen(port)
}

func (s *Server) HttpListen(port string) {
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("Failed to listen http")
	}
}

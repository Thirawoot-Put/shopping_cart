package server

import (
	"net/http"
	"os"

	"gorm.io/gorm"
)

type Server struct {
	mux *http.ServeMux
	db  *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	mux := http.NewServeMux()

	return &Server{mux: mux, db: db}
}

func (s *Server) Start() {
	port := os.Getenv("SERVER_PORT")

	s.initAuthRoute(s.db)
	s.initUserRoleRoute(s.db)

	s.HttpListen(port)
}

func (s *Server) HttpListen(port string) {
	err := http.ListenAndServe(":"+port, s.mux)
	if err != nil {
		panic(err)
	}
}

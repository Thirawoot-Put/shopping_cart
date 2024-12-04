package server

import (
	"Thirawoot/shopping_cart/internal/interfaces/handlers"
	"Thirawoot/shopping_cart/internal/repositories"
	"Thirawoot/shopping_cart/internal/services"
	"Thirawoot/shopping_cart/internal/usecase"
	"net/http"

	"gorm.io/gorm"
)

func (s *Server) initAuthRoute(db *gorm.DB) {
	authMux := http.NewServeMux()

	rp := repositories.NewAuthRepository()
	sv := services.NewAuthService(rp)
	uc := usecase.NewAuthUseCase(sv)
	hl := handlers.NewAuthHandler(uc)

	authMux.HandleFunc("/register", hl.PostUser)

	s.mux.Handle("/auth/", http.StripPrefix("/auth", authMux))
}

package server

import (
	"Thirawoot/shopping_cart/internal/interfaces/handlers"
	"Thirawoot/shopping_cart/internal/repositories"
	"Thirawoot/shopping_cart/internal/services"
	"Thirawoot/shopping_cart/internal/usecase"

	"gorm.io/gorm"
)

func (s *Server) initAuthRoute(db *gorm.DB) {
	rp := repositories.NewAuthRepository(db)
	sv := services.NewAuthService(rp)
	uc := usecase.NewAuthUseCase(sv)
	hl := handlers.NewAuthHandler(uc)

	s.mux.HandleFunc("/register/admin", hl.HandleAdmin)
	s.mux.HandleFunc("/register", hl.HandleCustomer)
	s.mux.HandleFunc("/login", hl.LoginUser)
}

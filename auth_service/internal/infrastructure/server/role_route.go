package server

import (
	"Thirawoot/shopping_cart/internal/interfaces/handlers"
	"Thirawoot/shopping_cart/internal/repositories"
	"Thirawoot/shopping_cart/internal/services"
	"Thirawoot/shopping_cart/internal/usecase"

	"gorm.io/gorm"
)

func (s *Server) initUserRoleRoute(db *gorm.DB) {
	rp := repositories.NewUserRoleRepository(db)
	sv := services.NewUserRoleService(rp)
	uc := usecase.NewUserRoleUseCasee(sv)
	hl := handlers.NewUserRoleHandler(*uc)

	s.mux.HandleFunc("/user-role", hl.HandleRole)
	s.mux.HandleFunc("/user-role/", hl.HandleRoleById)
}

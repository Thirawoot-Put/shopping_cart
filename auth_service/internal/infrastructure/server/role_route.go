package server

import (
	"Thirawoot/shopping_cart/internal/interfaces/handlers"
	"Thirawoot/shopping_cart/internal/repositories"
	"Thirawoot/shopping_cart/internal/services"
	"Thirawoot/shopping_cart/internal/usecase"
	"net/http"

	"gorm.io/gorm"
)

func (s *Server) initUserRoleRoute(db *gorm.DB) {
	roleMux := http.NewServeMux()

	rp := repositories.NewUserRoleRepository(db)
	sv := services.NewUserRoleService(rp)
	uc := usecase.NewUserRoleUseCasee(sv)
	hl := handlers.NewUserRoleHandler(*uc)

	roleMux.HandleFunc("/", hl.PostUserRole)

	s.mux.Handle("/user-role/", http.StripPrefix("/user-role", roleMux))
}

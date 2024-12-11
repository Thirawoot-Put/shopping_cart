package handlers

import (
	"Thirawoot/shopping_cart/internal/dto"
	"Thirawoot/shopping_cart/internal/usecase"
	"encoding/json"
	"net/http"
)

type UserRoleHandler struct {
	usecase usecase.UserRoleUseCase
}

func NewUserRoleHandler(u usecase.UserRoleUseCase) *UserRoleHandler {
	return &UserRoleHandler{usecase: u}
}

func (h *UserRoleHandler) PostUserRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	}

	var role dto.UserRoleCreate

	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	h.usecase.CreateRole(&role)
}

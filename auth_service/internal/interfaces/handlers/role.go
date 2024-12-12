package handlers

import (
	"Thirawoot/shopping_cart/internal/dto"
	"Thirawoot/shopping_cart/internal/shared/status"
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
		http.Error(w, "Method is not allowed", status.MethodNotAllowed)
	}

	var newRole dto.UserRoleCreate

	err := json.NewDecoder(r.Body).Decode(&newRole)
	if err != nil {
		http.Error(w, "Invalid json body", http.StatusBadRequest)
		return
	}

	res := h.usecase.CreateRole(&newRole)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encoding response", status.InternalError)
	}
}

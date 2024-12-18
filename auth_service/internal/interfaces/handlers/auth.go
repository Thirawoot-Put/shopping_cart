package handlers

import (
	"Thirawoot/shopping_cart/internal/dto"
	"Thirawoot/shopping_cart/internal/shared/status"
	"Thirawoot/shopping_cart/internal/usecase"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	usecase *usecase.AuthUseCase
}

func NewAuthHandler(u *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

func (h *AuthHandler) HandleAdmin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.postAdminUser(w, r)
	default:
		http.Error(w, "Not found method", status.MethodNotAllowed)
	}
}

func (h *AuthHandler) postAdminUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	}

	var newUser dto.UserCreate

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid json body", http.StatusBadRequest)
		return
	}

	res := h.usecase.CreateAdmin(&newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encoding response", status.InternalError)
	}
}

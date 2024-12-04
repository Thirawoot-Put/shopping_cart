package handlers

import (
	"Thirawoot/shopping_cart/internal/usecase"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	usecase *usecase.AuthUseCase
}

func NewAuthHandler(u *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

func (h *AuthHandler) PostUser(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	}

	fmt.Println("***> hello from handler")
	h.usecase.CreateUser()
}

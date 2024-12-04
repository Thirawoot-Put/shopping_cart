package handlers

import (
	"Thirawoot/shopping_cart/internal/usecase"
	"fmt"
)

type AuthHandler struct {
	usecase *usecase.AuthUseCase
}

func NewAuthHandler(u *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

func (h *AuthHandler) PostUser() {
	fmt.Println("***> hello from handler")
	h.usecase.CreateUser()
}

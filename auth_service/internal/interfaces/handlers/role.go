package handlers

import (
	"Thirawoot/shopping_cart/internal/dto"
	"Thirawoot/shopping_cart/internal/shared/status"
	"Thirawoot/shopping_cart/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type UserRoleHandler struct {
	usecase usecase.UserRoleUseCase
}

func NewUserRoleHandler(u usecase.UserRoleUseCase) *UserRoleHandler {
	return &UserRoleHandler{usecase: u}
}

func (h *UserRoleHandler) HandleRole(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getRoles(w, r)
	case http.MethodPost:
		h.postUserRole(w, r)
	default:
		http.Error(w, "Method is not allowed", status.MethodNotAllowed)
	}
}

func (h *UserRoleHandler) postUserRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", status.MethodNotAllowed)
		return
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

func (h *UserRoleHandler) getRoles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", status.MethodNotAllowed)
		return
	}

	res := h.usecase.FindRoles()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encoding response", status.InternalError)
	}
}

func (h *UserRoleHandler) HandleRoleById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getRole(w, r)
	case http.MethodDelete:
		h.deleteRole(w, r)
	default:
		http.Error(w, "Method is not allowed", status.MethodNotAllowed)
	}
}

func (h *UserRoleHandler) getRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", status.MethodNotAllowed)
		return
	}

	segments := strings.Split(r.URL.Path, "/")
	if segments[2] == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
	}

	id, err := strconv.Atoi(segments[2])
	if err != nil {
		http.Error(w, "ID must be integer", http.StatusBadRequest)
	}

	res := h.usecase.FindRole(uint(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encoding response", status.InternalError)
	}
}

func (h *UserRoleHandler) deleteRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method is not allowed", status.MethodNotAllowed)
		return
	}

	segments := strings.Split(r.URL.Path, "/")
	if segments[2] == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
	}

	id, err := strconv.Atoi(segments[2])
	if err != nil {
		http.Error(w, "ID must be integer", http.StatusBadRequest)
	}

	res := h.usecase.DeleteRole(uint(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encoding response", status.InternalError)
	}
}

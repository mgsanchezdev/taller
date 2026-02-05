package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"workshop-platform/backend/internal/auth"
	"workshop-platform/backend/internal/domain"
	"workshop-platform/backend/internal/repositories"
)

type AuthHandler struct {
	jwt    *auth.JWT
	userRepo repositories.UserRepository
}

func NewAuthHandler(jwt *auth.JWT, userRepo repositories.UserRepository) *AuthHandler {
	return &AuthHandler{jwt: jwt, userRepo: userRepo}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	WorkshopID string `json:"workshop_id"`
}

type LoginResponse struct {
	Token string     `json:"token"`
	User  domain.User `json:"user"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid body"}`, http.StatusBadRequest)
		return
	}
	if req.Email == "" || req.Password == "" || req.WorkshopID == "" {
		http.Error(w, `{"error":"email, password and workshop_id required"}`, http.StatusBadRequest)
		return
	}
	user, err := h.userRepo.GetByEmail(r.Context(), req.WorkshopID, req.Email)
	if err != nil || user == nil {
		http.Error(w, `{"error":"invalid credentials"}`, http.StatusUnauthorized)
		return
	}
	// TODO: compare password hash (bcrypt)
	// if !bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) { ... }

	token, err := h.jwt.Sign(user, 24*time.Hour)
	if err != nil {
		http.Error(w, `{"error":"token error"}`, http.StatusInternalServerError)
		return
	}
	user.PasswordHash = ""
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: token, User: *user})
}

package handlers

import (
	"net/http"

	"github.com/rmiguelac/logbook/backend/internal/repositories"
	"github.com/rmiguelac/logbook/backend/pkg/auth"
)

type AuthHandler struct {
	userRepo  *repositories.UserRepository
	jwtSecret string
}

func NewAuthHandler(userRepo *repositories.UserRepository, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := decodeJSONBody(w, r, &req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	// Get user from database
	user, err := h.userRepo.GetByEmail(r.Context(), req.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Verify password (using bcrypt)
	if err := auth.VerifyPassword(user.Password, req.Password); err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Generate JWT
	token, err := auth.GenerateJWT(user.ID, user.Email, h.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"token":   token,
		"user_id": user.ID.String(),
	})
}

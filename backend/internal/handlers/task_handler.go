package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rmiguelac/logbook/backend/internal/models"
	"github.com/rmiguelac/logbook/backend/internal/repositories"
	"github.com/rmiguelac/logbook/backend/pkg/auth"
)

type TaskHandler struct {
	repo *repositories.TaskRepository
}

func NewTaskHandler(repo *repositories.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

// CreateTaskRequest - Input validation struct
type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"oneof=todo ongoing halted done"`
}

// CreateTask - @Summary Create new task
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserIDFromContext(r.Context())
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Authentication required")
		return
	}
	var req CreateTaskRequest
	if err := decodeJSONBody(w, r, &req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Create task using repository
	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		UserID:      getUserIDFromContext(r.Context()), // From auth middleware
	}

	if err := h.repo.Create(r.Context(), task); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create task")
		return
	}

	respondWithJSON(w, http.StatusCreated, task)
}

// GetTask - @Summary Get task by ID
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := uuid.Parse(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	task, err := h.repo.GetByID(r.Context(), taskID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Task not found")
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

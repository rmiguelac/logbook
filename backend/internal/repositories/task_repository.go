package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/rmiguelac/logbook/backend/internal/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *TaskRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	var task models.Task
	err := r.db.WithContext(ctx).
		Preload("Tags").
		Preload("History").
		Preload("LinkedNotes").
		First(&task, "id = ?", id).
		Error
	return &task, err
}

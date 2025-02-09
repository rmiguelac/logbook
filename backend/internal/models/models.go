package models

import (
	"time"

	"github.com/google/uuid"
)

// Task - Main task entity
type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title       string    `gorm:"not null"`
	Description string
	Status      string `gorm:"not null;index"` // "todo", "ongoing", "halted", "done"
	DueDate     *time.Time
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	// Relationships
	UserID      uuid.UUID     `gorm:"type:uuid;not null;index"` // Task owner
	Tags        []Tag         `gorm:"many2many:task_tags;constraint:OnDelete:CASCADE"`
	History     []TaskHistory `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE"`
	LinkedNotes []Note        `gorm:"many2many:task_notes;constraint:OnDelete:CASCADE"`
}

// Comment - Individual diary-like entries on tasks
type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// Relationships
	TaskID uuid.UUID `gorm:"type:uuid;not null;index"`
	UserID uuid.UUID `gorm:"type:uuid;not null"` // Author
}

// TaskHistory - Status changes and system-generated events
type TaskHistory struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	TaskID    uuid.UUID `gorm:"type:uuid;not null;index"`
	EventType string    `gorm:"not null"` // "status_change", "comment", "note_link"
	OldStatus string    // Only populated for status changes
	NewStatus string
	CommentID *uuid.UUID `gorm:"type:uuid"` // Optional link to full comment
	NoteID    *uuid.UUID `gorm:"type:uuid"` // Optional link to note
	CreatedAt time.Time  `gorm:"autoCreateTime"`
}

// Note - Knowledge items/wiki entries
type Note struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relationships
	LinkedTasks []Task `gorm:"many2many:task_notes;constraint:OnDelete:CASCADE"`
	Tags        []Tag  `gorm:"many2many:note_tags;constraint:OnDelete:CASCADE"`
}

// Tag - Categorization system
type Tag struct {
	ID    uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name  string    `gorm:"uniqueIndex;not null"`
	Color string    `gorm:"not null"` // Hex code like "#FF0000"
}

// User - Authentication
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"` // bcrypt hash
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

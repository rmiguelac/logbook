package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type contextKey string

const (
	userIDKey contextKey = "userID"
)

// GetUserIDFromContext safely extracts user ID from request context
func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	val := ctx.Value(userIDKey)
	if val == nil {
		return uuid.Nil, fmt.Errorf("no user ID in context")
	}

	userID, ok := val.(uuid.UUID)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid user ID type in context")
	}

	return userID, nil
}

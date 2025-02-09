package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// GetUserIDFromContext safely extracts the user ID from context
func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	claims := GetUserFromContext(ctx)
	if claims == nil || claims.UserID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("user not authenticated")
	}
	return claims.UserID, nil
}

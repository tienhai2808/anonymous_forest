package repository

import (
	"context"

	"github.com/tienhai2808/anonymous_forest/backend/internal/model"
)

type PostRepository interface {
	Create(ctx context.Context, post *model.Post) error

	CountByClientID(ctx context.Context, clientID string) (int64, error)
}

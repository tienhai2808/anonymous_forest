package service

import (
	"context"

	"github.com/tienhai2808/anonymous_forest/backend/internal/request"
)

type PostService interface {
	CreatePost(ctx context.Context, clientID string, req request.CreatePostRequest) (string, error)
}
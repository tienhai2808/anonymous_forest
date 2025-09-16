package service

import (
	"context"

	"github.com/tienhai2808/anonymous_forest/backend/internal/request"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostService interface {
	CreatePost(ctx context.Context, clientID string, req request.CreatePostRequest) (string, error)

	GetPostByLink(ctx context.Context, postLink string) (bson.M, error)
}
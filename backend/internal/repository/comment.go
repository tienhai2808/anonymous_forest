package repository

import (
	"context"

	"github.com/tienhai2808/anonymous_forest/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *model.Comment) error

	DeleteByPostID(ctx context.Context, postID bson.ObjectID) error
}
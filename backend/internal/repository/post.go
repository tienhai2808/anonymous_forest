package repository

import (
	"context"

	"github.com/tienhai2808/anonymous_forest/backend/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostRepository interface {
	Create(ctx context.Context, post *model.Post) error

	FindByID(ctx context.Context, objID bson.ObjectID) (bson.M, error)

	FindRandomExcludeIDs(ctx context.Context, objIDs []bson.ObjectID) (bson.M, error)
}

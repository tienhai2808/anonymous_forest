package repository

import (
	"context"

	"github.com/tienhai2808/anonymous_forest/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostRepository interface {
	Create(ctx context.Context, post *model.Post) error

	FindByIDWithComments(ctx context.Context, objID bson.ObjectID) (bson.M, error)

	FindByID(ctx context.Context, objID bson.ObjectID) (*model.Post, error)

	Update(ctx context.Context, objID bson.ObjectID, data any) error

	FindRandomExcludeIDsWithComments(ctx context.Context, objIDs []bson.ObjectID) (bson.M, error)

	Delete(ctx context.Context, objID bson.ObjectID) error
}

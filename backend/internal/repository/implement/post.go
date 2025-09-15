package implement

import (
	"context"
	"time"

	"github.com/tienhai2808/anonymous_forest/backend/internal/common"
	"github.com/tienhai2808/anonymous_forest/backend/internal/model"
	"github.com/tienhai2808/anonymous_forest/backend/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type postRepositoryImpl struct {
	collection *mongo.Collection
}

func NewPostRepository(db *mongo.Database) repository.PostRepository {
	collection := db.Collection(common.PostCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "created_at", Value: -1}},
		},
		{
			Keys:    bson.D{{Key: "client_id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	collection.Indexes().CreateMany(ctx, indexes)

	return &postRepositoryImpl{collection}
}

func (r *postRepositoryImpl) Create(ctx context.Context, post *model.Post) error {
	if _, err := r.collection.InsertOne(ctx, post); err != nil {
		return err
	}

	return nil
}

func (r *postRepositoryImpl) CountByClientID(ctx context.Context, clientID string) (int64, error) {
	return r.collection.CountDocuments(ctx, bson.M{
		"client_id": clientID,
	})
}

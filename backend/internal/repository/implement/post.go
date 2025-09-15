package implement

import (
	"context"
	"time"

	"github.com/tienhai2808/anonymous_forest/backend/internal/common"
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
			Keys:    bson.D{{Key: "client_id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{{Key: "created_at", Value: -1}},
		},
	}

	collection.Indexes().CreateMany(ctx, indexes)

	return &postRepositoryImpl{collection}
}



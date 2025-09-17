package implement

import (
	"context"
	"time"

	"github.com/tienhai2808/anonymous_forest/internal/common"
	"github.com/tienhai2808/anonymous_forest/internal/model"
	"github.com/tienhai2808/anonymous_forest/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type commentRepositoryImpl struct {
	collection *mongo.Collection
}

func NewCommentRepository(db *mongo.Database) repository.CommentRepository {
	collection := db.Collection(common.CommentCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "created_at", Value: -1}},
		},
	}

	collection.Indexes().CreateMany(ctx, indexes)

	return &commentRepositoryImpl{collection}
}

func (r *commentRepositoryImpl) Create(ctx context.Context, comment *model.Comment) error {
	_, err := r.collection.InsertOne(ctx, comment)
	return err
}

func (r *commentRepositoryImpl) DeleteByPostID(ctx context.Context, postID bson.ObjectID) error {
	_, err := r.collection.DeleteMany(ctx, bson.M{"post_id": postID})
	return err
}

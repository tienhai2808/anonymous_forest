package implement

import (
	"context"
	"time"

	"github.com/tienhai2808/anonymous_forest/backend/internal/common"
	"github.com/tienhai2808/anonymous_forest/backend/internal/model"
	"github.com/tienhai2808/anonymous_forest/backend/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
			Keys: bson.D{{Key: "client_id", Value: 1}},
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

func (r *postRepositoryImpl) FindByID(ctx context.Context, objID bson.ObjectID) (bson.M, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "_id", Value: objID}}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "comments"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "post_id"},
			{Key: "as", Value: "comments"},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "comment_ids", Value: 0},
			{Key: "client_id", Value: 0},
			{Key: "_id", Value: 0},
			{Key: "updated_at", Value: 0},
		}}},
	}
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return results[0], nil
}

func (r *postRepositoryImpl) CountByClientID(ctx context.Context, clientID string) (int64, error) {
	return r.collection.CountDocuments(ctx, bson.M{
		"client_id": clientID,
	})
}

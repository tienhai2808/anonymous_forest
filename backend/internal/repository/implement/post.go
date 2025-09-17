package implement

import (
	"context"
	"errors"
	"time"

	"github.com/tienhai2808/anonymous_forest/internal/common"
	"github.com/tienhai2808/anonymous_forest/internal/model"
	"github.com/tienhai2808/anonymous_forest/internal/repository"
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
	}

	collection.Indexes().CreateMany(ctx, indexes)

	return &postRepositoryImpl{collection}
}

func (r *postRepositoryImpl) Create(ctx context.Context, post *model.Post) error {
	_, err := r.collection.InsertOne(ctx, post)
	return err
}

func (r *postRepositoryImpl) FindByIDWithComments(ctx context.Context, objID bson.ObjectID) (bson.M, error) {
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
			{Key: "updated_at", Value: 0},
		}}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return results[0], nil
}

func (r *postRepositoryImpl) FindRandomExcludeIDsWithComments(ctx context.Context, objIDs []bson.ObjectID) (bson.M, error) {
	match := bson.D{}
	if len(objIDs) > 0 {
		match = bson.D{{Key: "_id", Value: bson.D{{Key: "$nin", Value: objIDs}}}}
	}

	pipeline := mongo.Pipeline{}
	if len(match) > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: match}})
	}
	pipeline = append(pipeline, bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}},
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "comments"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "post_id"},
			{Key: "as", Value: "comments"},
		}}},
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "comment_ids", Value: 0},
			{Key: "updated_at", Value: 0},
		}}})

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return results[0], nil
}

func (r *postRepositoryImpl) Update(ctx context.Context, objID bson.ObjectID, updateData any) error {
	result, err := r.collection.UpdateByID(ctx, objID, updateData)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (r *postRepositoryImpl) FindByID(ctx context.Context, objID bson.ObjectID) (*model.Post, error) {
	var post model.Post
	if err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&post); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &post, nil
}

func (r *postRepositoryImpl) Delete(ctx context.Context, objID bson.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

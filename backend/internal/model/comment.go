package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Comment struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	PostID    bson.ObjectID `bson:"post_id" json:"post_id"`
	Content   string        `bson:"content" json:"content" validate:"required,min=2"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
}

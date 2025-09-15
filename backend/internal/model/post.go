package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Post struct {
	ID           bson.ObjectID   `bson:"_id,omitempty" json:"id"`
	ClientID     string          `bson:"client_id" json:"client_id"`
	Content      string          `bson:"content" json:"content"`
	EmpathyCount int             `bson:"empathy_count" json:"empathy_count"`
	ProtestCount int             `bson:"protest_count" json:"protest_count"`
	CommentIDs   []bson.ObjectID `bson:"comment_ids" json:"comment_ids"`
	CreatedAt    time.Time       `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time       `bson:"updated_at" json:"updated_at"`
}

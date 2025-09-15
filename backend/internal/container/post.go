package container

import (
	"github.com/tienhai2808/anonymous_forest/backend/internal/handler"
	repoImpl "github.com/tienhai2808/anonymous_forest/backend/internal/repository/implement"
	svcImpl "github.com/tienhai2808/anonymous_forest/backend/internal/service/implement"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PostContainer struct {
	PostHandler *handler.PostHandler
}

func NewPostContainer(db *mongo.Database) *PostContainer {
	postRepo := repoImpl.NewPostRepository(db)
	postSvc := svcImpl.NewPostService(postRepo)
	postHdl := handler.NewPostHandler(postSvc)

	return &PostContainer{postHdl}
}

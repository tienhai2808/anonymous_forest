package implement

import (
	"github.com/tienhai2808/anonymous_forest/backend/internal/repository"
	"github.com/tienhai2808/anonymous_forest/backend/internal/service"
)

type postServiceImpl struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) service.PostService {
	return &postServiceImpl{postRepo}
}

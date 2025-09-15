package implement

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/tienhai2808/anonymous_forest/backend/internal/common"
	"github.com/tienhai2808/anonymous_forest/backend/internal/model"
	"github.com/tienhai2808/anonymous_forest/backend/internal/repository"
	"github.com/tienhai2808/anonymous_forest/backend/internal/request"
	"github.com/tienhai2808/anonymous_forest/backend/internal/service"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type postServiceImpl struct {
	postRepo  repository.PostRepository
	redisRepo repository.RedisRepository
}

func NewPostService(postRepo repository.PostRepository, redisRepo repository.RedisRepository) service.PostService {
	return &postServiceImpl{
		postRepo,
		redisRepo,
	}
}

func (s *postServiceImpl) CreatePost(ctx context.Context, clientID string, req request.CreatePostRequest) (string, error) {
	postCount, err := s.postRepo.CountByClientID(ctx, clientID)
	if err != nil {
		return "", fmt.Errorf("kiểm tra số lượng bài viết thất bại: %w", err)
	}
	if postCount >= 5 {
		return "", common.ErrTooManyPost
	}

	post := &model.Post{
		ID:           bson.NewObjectID(),
		ClientID:     clientID,
		Content:      req.Content,
		EmpathyCount: 0,
		ProtestCount: 0,
		CommentIDs:   []bson.ObjectID{},
		CreatedAt:    time.Now().Local(),
		UpdatedAt:    time.Now().Local(),
	}

	if err := s.postRepo.Create(ctx, post); err != nil {
		return "", fmt.Errorf("tạo bài viết thất bại: %w", err)
	}

	var postLink string
	if req.GetLink != nil && *req.GetLink {
		postLink = randomLink()
		key := fmt.Sprintf("get-link:%s", postLink)
		if err := s.redisRepo.SaveString(ctx, key, post.ID.String(), 7*24*time.Hour); err != nil {
			return "", fmt.Errorf("tạo liên kết theo dõi thất bại: %w", err)
		}
	}

	return postLink, nil
}

func randomLink() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, 6)
	for i := range result {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[n.Int64()]
	}

	return string(result)
}

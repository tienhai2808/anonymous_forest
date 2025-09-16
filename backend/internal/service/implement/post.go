package implement

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/tienhai2808/anonymous_forest/backend/internal/common"
	"github.com/tienhai2808/anonymous_forest/backend/internal/model"
	"github.com/tienhai2808/anonymous_forest/backend/internal/repository"
	"github.com/tienhai2808/anonymous_forest/backend/internal/request"
	"github.com/tienhai2808/anonymous_forest/backend/internal/service"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
	key := fmt.Sprintf("post-created:%s", clientID)
	postCount, err := s.redisRepo.IncrementWithTTL(ctx, key, 24*time.Hour)
	if err != nil {
		return "", fmt.Errorf("tăng số lượng bài viết thất bại: %w", err)
	}
	if postCount > 5 {
		if err = s.redisRepo.Decrement(ctx, key); err != nil {
			return "", fmt.Errorf("giảm số lượng bài viết thất bại: %w", err)
		}
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
		if err = s.redisRepo.Decrement(ctx, key); err != nil {
			return "", fmt.Errorf("giảm số lượng bài viết thất bại: %w", err)
		}
		return "", fmt.Errorf("tạo bài viết thất bại: %w", err)
	}

	var postLink string
	if req.GetLink != nil && *req.GetLink {
		postLink = randomLink()
		key = fmt.Sprintf("get-link:%s", postLink)
		if err := s.redisRepo.SetString(ctx, key, post.ID.Hex(), 7*24*time.Hour); err != nil {
			if err = s.redisRepo.Decrement(ctx, key); err != nil {
				return "", fmt.Errorf("giảm số lượng bài viết thất bại: %w", err)
			}
			return "", fmt.Errorf("tạo liên kết theo dõi thất bại: %w", err)
		}
	}

	return postLink, nil
}

func (s *postServiceImpl) GetPostByLink(ctx context.Context, postLink string) (bson.M, error) {
	key := fmt.Sprintf("get-link:%s", postLink)
	postID, err := s.redisRepo.GetString(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("lấy thông tin bài viết thất bại: %w", err)
	}
	if postID == "" {
		return nil, common.ErrPostNotFound
	}

	postObjID, err := bson.ObjectIDFromHex(postID)
	if err != nil {
		return nil, fmt.Errorf("ID không hợp lệ: %w", err)
	}

	post, err := s.postRepo.FindByID(ctx, postObjID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.ErrPostNotFound
		}
		return nil, fmt.Errorf("lấy thông tin bài viết thất bại: %w", err)
	}

	return post, nil
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

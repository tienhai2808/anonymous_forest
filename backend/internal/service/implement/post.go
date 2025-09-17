package implement

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/tienhai2808/anonymous_forest/internal/common"
	"github.com/tienhai2808/anonymous_forest/internal/model"
	"github.com/tienhai2808/anonymous_forest/internal/repository"
	"github.com/tienhai2808/anonymous_forest/internal/request"
	"github.com/tienhai2808/anonymous_forest/internal/service"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type postServiceImpl struct {
	postRepo  repository.PostRepository
	cmtRepo   repository.CommentRepository
	redisRepo repository.RedisRepository
}

func NewPostService(postRepo repository.PostRepository, cmtRepo repository.CommentRepository, redisRepo repository.RedisRepository) service.PostService {
	return &postServiceImpl{
		postRepo,
		cmtRepo,
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
		return "", common.ErrTooManyPostsCreated
	}

	post := &model.Post{
		ID:           bson.NewObjectID(),
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
		return nil, fmt.Errorf("chuyển đổi ID thành ObjectID thất bại: %w", err)
	}

	post, err := s.postRepo.FindByIDWithComments(ctx, postObjID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.ErrPostNotFound
		}
		return nil, fmt.Errorf("lấy thông tin bài viết thất bại: %w", err)
	}

	return post, nil
}

func (s *postServiceImpl) GetRandomPost(ctx context.Context, clientID string) (bson.M, error) {
	key := fmt.Sprintf("viewed-posts:%s", clientID)
	viewedIDs, err := s.redisRepo.SetMembers(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("lấy danh sách bài viết đã xem thất bại: %w", err)
	}

	if len(viewedIDs) >= 10 {
		return nil, common.ErrTooManyPostsViewed
	}

	objectIDs := make([]bson.ObjectID, 0, len(viewedIDs))
	for _, id := range viewedIDs {
		objectID, err := bson.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf("chuyển đổi ID thành ObjectID thất bại: %w", err)
		}
		objectIDs = append(objectIDs, objectID)
	}

	post, err := s.postRepo.FindRandomExcludeIDsWithComments(ctx, objectIDs)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.ErrPostNotFound
		}
		return nil, fmt.Errorf("lấy ngẫu nhiên bài viết thất bại: %w", err)
	}

	if oid, ok := post["_id"].(bson.ObjectID); ok {
		if err = s.redisRepo.SetAddWithTTL(ctx, key, oid.Hex(), 24*time.Hour); err != nil {
			return nil, fmt.Errorf("lưu bài viết đã xem thất bại: %w", err)
		}
	}

	return post, nil
}

func (s *postServiceImpl) AddEmpathyPost(ctx context.Context, id string) error {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return common.ErrInvalidID
	}

	updateData := bson.M{
		"$inc": bson.M{"empathy_count": 1},
		"$set": bson.M{"updated_at": time.Now().Local()},
	}
	if err = s.postRepo.Update(ctx, oid, updateData); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return common.ErrPostNotFound
		}
		return fmt.Errorf("tăng số lượng đồng cảm thất bại: %w", err)
	}

	return nil
}

func (s *postServiceImpl) AddProtestPost(ctx context.Context, id string) error {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return common.ErrInvalidID
	}

	updateData := bson.M{
		"$inc": bson.M{"protest_count": 1},
		"$set": bson.M{"updated_at": time.Now().Local()},
	}
	if err = s.postRepo.Update(ctx, oid, updateData); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return common.ErrPostNotFound
		}
		return fmt.Errorf("tăng số lượng phản đối bài viết thất bại: %w", err)
	}

	post, err := s.postRepo.FindByID(ctx, oid)
	if err != nil {
		return fmt.Errorf("lấy thông tin bài viết thất bại: %w", err)
	}
	if post == nil {
		return common.ErrPostNotFound
	}

	if post.ProtestCount >= 5 {
		if err = s.postRepo.Delete(ctx, oid); err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return common.ErrPostNotFound
			}
			return fmt.Errorf("xóa bài viết thất bại: %w", err)
		}

		if err = s.cmtRepo.DeleteByPostID(ctx, oid); err != nil {
			return fmt.Errorf("xóa bình luận của bài viết thất bại: %w", err)
		}
	}

	return nil
}

func (s *postServiceImpl) CreatePostComment(ctx context.Context, id string, req request.CreatePostCommentRequest) error {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return common.ErrInvalidID
	}

	post, err := s.postRepo.FindByID(ctx, oid)
	if err != nil {
		return fmt.Errorf("lấy thông tin bài viết thất bại: %w", err)
	}
	if post == nil {
		return common.ErrPostNotFound
	}

	comment := &model.Comment{
		ID:        bson.NewObjectID(),
		PostID:    oid,
		Content:   req.Content,
		CreatedAt: time.Now().Local(),
	}

	if err = s.cmtRepo.Create(ctx, comment); err != nil {
		return fmt.Errorf("thêm bình luận thất bại: %w", err)
	}

	updateData := bson.M{
		"$push": bson.M{"comment_ids": comment.ID},
		"$set":  bson.M{"updated_at": time.Now().Local()},
	}
	if err = s.postRepo.Update(ctx, oid, updateData); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return common.ErrPostNotFound
		}
		return fmt.Errorf("tăng số lượng phản đối bài viết thất bại: %w", err)
	}

	return nil
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

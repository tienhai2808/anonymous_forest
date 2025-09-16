package handler

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tienhai2808/anonymous_forest/backend/internal/common"
	"github.com/tienhai2808/anonymous_forest/backend/internal/request"
	"github.com/tienhai2808/anonymous_forest/backend/internal/service"
)

var validate = validator.New()

type PostHandler struct {
	postSvc service.PostService
}

func NewPostHandler(postSvc service.PostService) *PostHandler {
	return &PostHandler{postSvc}
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	clientID := c.Locals("client_id").(string)

	var req request.CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		message := common.HandleValidationError(err)
		return common.JSON(c, fiber.StatusBadRequest, message, nil)
	}

	if err := validate.Struct(req); err != nil {
		message := common.HandleValidationError(err)
		return common.JSON(c, fiber.StatusBadRequest, message, nil)
	}

	postLink, err := h.postSvc.CreatePost(ctx, clientID, req)
	if err != nil {
		switch err {
		case common.ErrTooManyPost:
			return common.JSON(c, fiber.StatusTooManyRequests, err.Error(), nil)
		default:
			return common.JSON(c, fiber.StatusInternalServerError, err.Error(), nil)
		}
	}

	return common.JSON(c, fiber.StatusCreated, "Tạo bài viết thành công", fiber.Map{
		"post_link": postLink,
	})
}

func (h *PostHandler) GetPostByLink(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	postLink := c.Params("id")

	post, err := h.postSvc.GetPostByLink(ctx, postLink)
	if err != nil {
		switch err {
		case common.ErrPostNotFound:
			return common.JSON(c, fiber.StatusNotFound, err.Error(), nil)
		default:
			return common.JSON(c, fiber.StatusInternalServerError, err.Error(), nil)
		}
	}

	return common.JSON(c, fiber.StatusOK, "Lấy bài viết thành công", fiber.Map{
		"post": post,
	})
}

// func (h *PostHandler) GetRandomPost(c *fiber.Ctx) error {
// 	_, cancel := context.WithTimeout(c.Context(), 5*time.Second)
// 	defer cancel()

// 	return common.JSON(c, fiber.StatusOK, "Lấy bài viết thành công", nil)
// }

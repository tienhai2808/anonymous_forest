package handler

import (
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
	var req request.CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		message := common.HandleValidationError(err)
		return common.JSON(c, fiber.StatusBadRequest, message, nil)
	}

	if err := validate.Struct(req); err != nil {
		message := common.HandleValidationError(err)
		return common.JSON(c, fiber.StatusBadRequest, message, nil)
	}

	return common.JSON(c, fiber.StatusCreated, "Tạo bài viết thành công", fiber.Map{
		"post_id": 1,
	})
}

package request

type CreatePostRequest struct {
	Content string `json:"content" validate:"required,min=2"`
	GetLink *bool  `json:"get_link" validate:"required"`
}

type CreatePostCommentRequest struct {
	Content string `json:"content" validate:"required,min=2"`
}

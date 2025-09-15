package request

type CreatePostRequest struct {
	Content      string `json:"content" validate:"required,min=2"`
	EmpathyCount int    `json:"empathy_count" validate:"required,gte=0"`
}

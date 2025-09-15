package common

type ApiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

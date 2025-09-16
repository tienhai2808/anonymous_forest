package common

import "errors"

var (
	ErrTooManyPost = errors.New("bạn đã hết lượt tạo bài viết trong hôm nay")

	ErrPostNotFound = errors.New("không tìm thấy bài viết")
)
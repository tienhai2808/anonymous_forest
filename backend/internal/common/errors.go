package common

import "errors"

var (
	ErrTooManyPost = errors.New("bạn đã hết lượt tạo bài viết trong hôm nay")
)
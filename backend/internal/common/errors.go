package common

import "errors"

var (
	ErrTooManyPostsCreated = errors.New("bạn đã hết lượt tạo bài viết trong hôm nay")

	ErrTooManyPostsViewed = errors.New("bạn đã hết lượt xem bài viết trong hôm nay")

	ErrPostNotFound = errors.New("không tìm thấy bài viết")

	ErrInvalidID = errors.New("ID không hợp lệ")
)

package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tienhai2808/anonymous_forest/internal/handler"
)

func PostRouter(rg fiber.Router, postHdl *handler.PostHandler) {
	post := rg.Group("/posts")
	{
		post.Post("", postHdl.CreatePost)
		post.Get("", postHdl.GetRandomPost)
		post.Get("/:link", postHdl.GetPostByLink)
		post.Post("/:id/comments", postHdl.CreatePostComment)
		post.Patch("/:id/empathy", postHdl.AddEmpathyPost)
		post.Patch("/:id/protest", postHdl.AddProtestPost)
	}
}
package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tienhai2808/anonymous_forest/backend/internal/handler"
)

func PostRouter(rg fiber.Router, postHdl *handler.PostHandler) {
	post := rg.Group("/posts")
	{
		post.Post("", postHdl.CreatePost)
		// post.Get("", postHdl.GetRandomPost)
		post.Get("/:id", postHdl.GetPostByLink)
	}
}
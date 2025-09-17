package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tienhai2808/anonymous_forest/config"
)

func CheckSession(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientID := c.Cookies(cfg.App.ClientToken)
		if !isValidUUID(clientID) || clientID == "" {
			clientID = uuid.NewString()
			c.Cookie(&fiber.Cookie{
				Name:     cfg.App.ClientToken,
				Value:    clientID,
				Expires:  time.Now().Add(cfg.App.TokenExpiresIn * time.Hour),
				HTTPOnly: cfg.App.HttpCookie,
				Secure:   cfg.App.SecureCookie,
				Path:     "/",
			})
		}

		c.Locals("client_id", clientID)

		return c.Next()
	}
}

func isValidUUID(str string) bool {
	_, err := uuid.Parse(str)
	return err == nil
}

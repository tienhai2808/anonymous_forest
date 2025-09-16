package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/redis/go-redis/v9"
	"github.com/tienhai2808/anonymous_forest/backend/config"
	"github.com/tienhai2808/anonymous_forest/backend/internal/common"
	"github.com/tienhai2808/anonymous_forest/backend/internal/container"
	"github.com/tienhai2808/anonymous_forest/backend/internal/initialization"
	"github.com/tienhai2808/anonymous_forest/backend/internal/middleware"
	"github.com/tienhai2808/anonymous_forest/backend/internal/router"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	cfg *config.Config
	app *fiber.App
	mdb *mongo.Client
	rdb *redis.Client
}

func NewServer(cfg *config.Config) (*Server, error) {
	app := fiber.New(
		fiber.Config{
			Prefork:      cfg.App.Http.Prefork,
			WriteTimeout: cfg.App.Http.WriteTimeout * time.Second,
			ReadTimeout:  cfg.App.Http.ReadTimeout * time.Second,
			IdleTimeout:  cfg.App.Http.WriteTimeout * time.Second,
			BodyLimit:    cfg.App.Http.BodyLimit * 1024 * 1024,
		},
	)

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.App.Cors.AllowOrigins,
		AllowMethods:     cfg.App.Cors.AllowMethods,
		AllowHeaders:     cfg.App.Cors.AllowHeaders,
		ExposeHeaders:    cfg.App.Cors.ExposeHeaders,
		AllowCredentials: cfg.App.Cors.AllowCredentials,
		MaxAge:           cfg.App.Cors.MaxAge,
	}))

	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(common.ApiResponse{
				Message: "to many requests!!!",
				Data:    nil,
			})
		},
	}))

	app.Use(middleware.CheckSession(cfg))

	mdb, err := initialization.InitDB(cfg)
	if err != nil {
		return nil, err
	}

	rdb, err := initialization.InitCache(cfg)
	if err != nil {
		return nil, err
	}

	postCtn := container.NewPostContainer(mdb.Database(cfg.Database.DBName), rdb)

	api := app.Group(cfg.App.ApiPrefix)
	router.PostRouter(api, postCtn.PostHandler)

	return &Server{
		cfg,
		app,
		mdb,
		rdb,
	}, nil
}

func (s *Server) Start() error {
	return s.app.Listen(fmt.Sprintf(":%d", s.cfg.App.Port))
}

func (s *Server) Shutdown(ctx context.Context) {
	if !fiber.IsChild() {
		log.Println("Đang shutdown server...")
	}

	if s.mdb != nil {
		_ = s.mdb.Disconnect(ctx)
	}

	if s.rdb != nil {
		_ = s.rdb.Close()
	}

	if s.app != nil {
		_ = s.app.ShutdownWithContext(ctx)
	}

	if !fiber.IsChild() {
		log.Println("Shutdown server thành công")
	}
}

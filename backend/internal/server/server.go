package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tienhai2808/anonymous_forest/backend/config"
	"github.com/tienhai2808/anonymous_forest/backend/internal/container"
	"github.com/tienhai2808/anonymous_forest/backend/internal/initialization"
	"github.com/tienhai2808/anonymous_forest/backend/internal/router"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	cfg *config.Config
	app *fiber.App
	mCli  *mongo.Client
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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.App.Cors.AllowOrigins,
		AllowMethods:     cfg.App.Cors.AllowMethods,
		AllowHeaders:     cfg.App.Cors.AllowHeaders,
		ExposeHeaders:    cfg.App.Cors.ExposeHeaders,
		AllowCredentials: cfg.App.Cors.AllowCredentials,
		MaxAge:           cfg.App.Cors.MaxAge,
	}))

	mCli, err := initialization.InitDB(cfg)
	if err != nil {
		return nil, err
	}

	postCtn := container.NewPostContainer(mCli.Database(cfg.Database.DBName))

	api := app.Group(cfg.App.ApiPrefix)
	router.PostRouter(api, postCtn.PostHandler)

	return &Server{
		cfg,
		app,
		mCli,
	}, nil
}

func (s *Server) Start() error {
	return s.app.Listen(fmt.Sprintf(":%d", s.cfg.App.Port))
}

func (s *Server) Shutdown(ctx context.Context) {
	if !fiber.IsChild() {
		log.Println("Đang shutdown server...")
	}

	if s.mCli != nil {
		_ = s.mCli.Disconnect(ctx)
	}

	if s.app != nil {
		_ = s.app.ShutdownWithContext(ctx)
	}

	if !fiber.IsChild() {
		log.Println("Shutdown server thành công")
	}
}

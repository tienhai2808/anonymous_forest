package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tienhai2808/anonymous_forest/backend/config"
	"github.com/tienhai2808/anonymous_forest/backend/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Tải cấu hình server thất bại: %v", err)
	}

	server, err := server.NewServer(cfg)
	if err != nil {
		if !fiber.IsChild() {
			log.Fatalf("Khởi tạo server thất bại: %v", err)
		}
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	errCh := make(chan error, 1)

	go func() {
		if err := server.Start(); err != nil {
			errCh <- err
		}
	}()

	if !fiber.IsChild() {
		log.Println("Chạy server thành công")
	}

	select {
	case err = <-errCh:
		if !fiber.IsChild() {
			log.Printf("Chạy server thất bại: %v", err)
		}
	case <-stop:
		if !fiber.IsChild() {
			log.Println("Có tín hiệu dừng server")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}

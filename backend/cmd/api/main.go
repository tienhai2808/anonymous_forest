package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tienhai2808/anonymous_forest/internal/config"
	"github.com/tienhai2808/anonymous_forest/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Tải cấu hình server thất bại: %v", err)
	}

	s, err := server.NewServer(cfg)
	if err != nil {
		if !fiber.IsChild() {
			log.Fatalf("Khởi tạo server thất bại: %v", err)
		}
	}

	ch := make(chan error, 1)

	go func() {
		if err := s.Start(); err != nil {
			ch <- err
		}
	}()

	if !fiber.IsChild() {
		log.Println("Chạy server thành công")
	}

	s.GracefulShutdown(ch)
}

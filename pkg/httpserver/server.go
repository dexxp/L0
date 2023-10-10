package server

import (
	"fmt"

	"github.com/dexxp/L0/config"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Fiber *fiber.App
	Cfg *config.HTTP
}

func NewServer(cfg *config.HTTP) *Server {
	return &Server{
		Fiber: fiber.New(fiber.Config{DisableStartupMessage: true}),
		Cfg: cfg,
	}
}

func (s *Server) Run() error {
	fmt.Println("Server running...", s.Cfg.Host, s.Cfg.Port)
	err := s.Fiber.Listen(fmt.Sprintf("%s:%s", s.Cfg.Host, s.Cfg.Port))
	if err != nil {
		fmt.Printf("Cannot listen. Error: {%s}\n", err)
	}
	return nil
}

func (s *Server) HomeRouter() {
	s.Fiber.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, Fiber!")
	})
}
package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
	"time"
)

type Server struct {
	httpServer *fiber.App
}

func NewServer() *Server {
	return &Server{httpServer: fiber.New(fiber.Config{
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	})}
}

func (s *Server) Run(port string, handler fasthttp.RequestHandler) error {
	s.httpServer.Server().Handler = handler
	return s.httpServer.Listen(":" + port)
}

func (s *Server) Shutdown() error {
	return s.httpServer.Shutdown()
}

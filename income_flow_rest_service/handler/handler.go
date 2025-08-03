package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() fasthttp.RequestHandler {
	router := fiber.New()

	// auth
	auth := router.Group("/good")
	auth.Post("/get", h.GetOneGood)
	auth.Post("/get_all", h.GetAllGoods)
	auth.Post("/create", h.CreateGood)
	auth.Post("/update", h.UpdateGood)
	auth.Post("/delete", h.DeleteGood)

	return router.Handler()
}

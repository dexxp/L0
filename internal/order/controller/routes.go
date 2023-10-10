package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetOrderRoute(app *fiber.App, h *OrderHandler) {
	app.Get("/orders", h.getAllOrdersHandler)
	app.Get("/orders/:orderUid", h.getOrderHandler)
}


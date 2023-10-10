package controller

import (
	"github.com/dexxp/L0/internal/order/usecase"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	UC *usecase.OrderUseCase 
}

func NewOrderHandler(uc *usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		UC: uc,
	}
}

func (h *OrderHandler) getOrderHandler(c *fiber.Ctx) error {
	orderUid := c.Params("orderUid")

	order := h.UC.GetOrderFromCacheByID(orderUid)

	return c.JSON(order)
}

func (h *OrderHandler) getAllOrdersHandler(c *fiber.Ctx) error {
	orders := h.UC.GetAllOrdersFromCache()

	return c.JSON(orders)
}
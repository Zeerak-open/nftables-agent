package handlers

import (
	"nft-http-api/internal/transport/httpServer/fiber/payloads"

	"github.com/gofiber/fiber/v2"
)

func (h *handlers) Health(c *fiber.Ctx) error {
	return c.JSON(payloads.ErrorPayload{
		Status: true,
	})
}

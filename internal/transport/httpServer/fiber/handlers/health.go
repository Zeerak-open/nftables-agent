package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/logicalangel/Zeerak-nft-agent/internal/transport/httpServer/fiber/payloads"
)

func (h *handlers) Health(c *fiber.Ctx) error {
	return c.JSON(payloads.ErrorPayload{
		Status: true,
	})
}

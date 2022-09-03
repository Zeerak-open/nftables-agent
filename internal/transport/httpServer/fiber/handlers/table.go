package handlers

import (
	"fmt"
	"nft-http-api/internal/transport/httpServer/fiber/payloads"

	"github.com/gofiber/fiber/v2"
)

func (h *handlers) FindTables(c *fiber.Ctx) error {
	tables, err := h.nft.GetTables()
	if err != nil {
		return c.Status(500).JSON(payloads.ErrorPayload{
			Status:  false,
			Message: "Failed to fetch tables",
			Extra:   err.Error(),
		})
	}

	return c.JSON(tables)
}

func (h *handlers) FindTableByName(c *fiber.Ctx) error {
	tableName := c.Params("tableName", "")

	table, err := h.nft.GetTable(tableName)
	if err != nil {
		return c.Status(500).JSON(payloads.ErrorPayload{
			Status:  false,
			Message: "Failed to fetch table",
		})
	}

	if table == nil {
		return c.Status(404).JSON(payloads.ErrorPayload{
			Status:  false,
			Message: fmt.Sprintf("table %s notfound", tableName),
		})
	}

	return c.JSON(table)
}

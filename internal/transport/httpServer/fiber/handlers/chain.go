package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/logicalangel/Zeerak-nft-agent/internal/transport/httpServer/fiber/payloads"
)

func (h *handlers) FindChains(c *fiber.Ctx) error {
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

	chains, err := h.nft.GetChains(table)
	if err != nil {
		return c.Status(500).JSON(payloads.ErrorPayload{
			Status:  false,
			Message: "Failed to fetch chains",
			Extra:   err.Error(),
		})
	}

	return c.JSON(chains)
}

func (h *handlers) FindChainByName(c *fiber.Ctx) error {
	tableName := c.Params("tableName", "")
	chainName := c.Params("chainName", "")

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

	chain, err := h.nft.GetChain(table, chainName)
	if err != nil {
		return c.Status(500).JSON(payloads.ErrorPayload{
			Status:  false,
			Message: "Failed to fetch chain",
			Extra:   err.Error(),
		})
	}

	if chain == nil {
		return c.Status(500).JSON(payloads.ErrorPayload{
			Status:  false,
			Message: fmt.Sprintf("chain %s notfound", chainName),
		})
	}

	return c.JSON(chain)
}

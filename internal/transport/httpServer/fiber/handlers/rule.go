package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/logicalangel/Zeerak-nft-agent/internal/transport/httpServer/fiber/payloads"
)

func (h *handlers) FindRules(c *fiber.Ctx) error {
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

	rules, err := h.nft.GetRules(table, chain)
	if err != nil {
		return c.Status(500).JSON(payloads.ErrorPayload{
			Status:  false,
			Message: "Failed to fetch rules",
			Extra:   err.Error(),
		})
	}

	return c.JSON(rules)
}

func (h *handlers) FindRuleByID(c *fiber.Ctx) error {
	tableName := c.Params("tableName", "")
	chainName := c.Params("chainName", "")
	_ruleId := c.Params("ruleId", "")

	ruleId, err := strconv.ParseUint(_ruleId, 10, 64)
	if err != nil {
		return c.Status(500).JSON(payloads.ErrorPayload{
			Status:  false,
			Message: "Invalid rule ID",
			Extra:   err.Error(),
		})
	}

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

	rule, err := h.nft.GetRule(table, chain, ruleId)
	if err != nil {
		return c.Status(404).JSON(payloads.ErrorPayload{
			Status:  false,
			Message: "field to fetch rules",
			Extra:   err.Error(),
		})
	}

	if rule == nil {
		return c.Status(404).JSON(payloads.ErrorPayload{
			Status:  false,
			Message: fmt.Sprintf("rule %s notfound", _ruleId),
			Extra:   err.Error(),
		})
	}

	return c.JSON(rule)
}

package handlers

import (
	"nft-http-api/internal/transport/nft"

	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	Health(c *fiber.Ctx) error

	FindChains(c *fiber.Ctx) error
	FindChainByName(c *fiber.Ctx) error

	FindRules(c *fiber.Ctx) error
	FindRuleByID(c *fiber.Ctx) error

	FindTables(c *fiber.Ctx) error
	FindTableByName(c *fiber.Ctx) error
}

type handlers struct {
	nft *nft.NftService
}

func New(nft *nft.NftService) Handlers {
	return &handlers{nft: nft}
}

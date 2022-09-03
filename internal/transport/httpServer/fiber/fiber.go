package fiber

import (
	"nft-http-api/internal/transport/httpServer"
	"nft-http-api/internal/transport/httpServer/fiber/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberHttpServer struct {
	app      *fiber.App
	address  string
	handlers handlers.Handlers
}

// @title Zeerak nft-agent
// @version 1.0
// @description Zeerak api interface of nft-agent
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @BasePath /
func New(address string, handlers handlers.Handlers) httpServer.HttpServer {
	if address == "" {
		address = "0.0.0.0:3000"
	}
	return &FiberHttpServer{
		app: fiber.New(fiber.Config{
			AppName:               "Zeerak",
			ServerHeader:          "Zeerak",
			DisableStartupMessage: true,
		}),
		address:  address,
		handlers: handlers,
	}
}

func (fhs *FiberHttpServer) Listen() error {
	fhs.app.Use(cors.New())

	if err := fhs.app.Listen(fhs.address); err != nil {
		return err
	}
	return nil
}

func (fhs *FiberHttpServer) RegisterRoutes() {
	fhs.app.Get("/health", fhs.handlers.Health)

	fhs.app.Get("/table", fhs.handlers.FindTables)
	fhs.app.Get("/table/:tableName", fhs.handlers.FindTableByName)
	fhs.app.Get("/table/:tableName/chain", fhs.handlers.FindChains)
	fhs.app.Get("/table/:tableName/chain/:chainName", fhs.handlers.FindChainByName)
	fhs.app.Get("/table/:tableName/chain/:chainName/rule", fhs.handlers.FindRules)
	fhs.app.Get("/table/:tableName/chain/:chainName/rule/:ruleId", fhs.handlers.FindRuleByID)
}

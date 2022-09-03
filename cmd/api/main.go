package main

import (
	"log"

	"github.com/logicalangel/Zeerak-nft-agent/internal/transport/httpServer/fiber"
	"github.com/logicalangel/Zeerak-nft-agent/internal/transport/httpServer/fiber/handlers"
	"github.com/logicalangel/Zeerak-nft-agent/internal/transport/nft"
	"github.com/logicalangel/Zeerak-nft-agent/internal/utils"
)

func main() {
	_ = utils.LoadCertificates("", "")

	nft := nft.New()

	handlers := handlers.New(nft)

	httpServer := fiber.New("0.0.0.0:3000", handlers)
	if err := httpServer.Listen(); err != nil {
		log.Fatal(err)
	}
}

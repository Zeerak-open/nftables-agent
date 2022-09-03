package main

import (
	"log"
	"nft-http-api/internal/transport/httpServer/fiber"
	"nft-http-api/internal/transport/httpServer/fiber/handlers"
	"nft-http-api/internal/transport/nft"
	"nft-http-api/internal/utils"
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

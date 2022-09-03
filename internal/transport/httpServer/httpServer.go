package httpServer

type HttpServer interface {
	Listen() error
	RegisterRoutes()
}

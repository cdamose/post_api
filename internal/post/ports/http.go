package ports

import (
	"post_api/internal/post/container"
)

type HttpServer struct {
	Application container.Application
}

func NewHttpServer(application container.Application) HttpServer {

	return HttpServer{Application: application}
}

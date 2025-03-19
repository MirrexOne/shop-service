package v1

import (
	"net/http"
	"shop-service/internal/service"
)

type Router struct {
	Mux *http.ServeMux
}

func New() *Router {
	return &Router{
		Mux: http.DefaultServeMux,
	}
}

func NewRouter(router *Router, services *service.Services) {
	newAuthRoutes(router, services.Auth)
}

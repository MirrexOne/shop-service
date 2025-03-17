package v1

import (
	"net/http"
	"shop-service/internal/service"
)

type Router struct {
	Mux      *http.ServeMux
	Handlers map[string]http.Handler
}

func New() *Router {
	return &Router{
		Mux:      http.DefaultServeMux,
		Handlers: make(map[string]http.Handler),
	}
}

func NewRouter(router *Router, services *service.Services) {
	newAuthRoutes(router, services.Auth)
}

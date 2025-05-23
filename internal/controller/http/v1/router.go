package v1

import (
	"github.com/gorilla/mux"
	"shop-service/internal/service"
)

type Router struct {
	Mux        *mux.Router
	Middleware AuthMiddleware
}

func New() *Router {
	return &Router{
		Mux:        mux.NewRouter(),
		Middleware: AuthMiddleware{},
	}
}

func NewRouter(router *Router, services *service.Services) {
	newAuthRoutes(router, services.Auth)
}

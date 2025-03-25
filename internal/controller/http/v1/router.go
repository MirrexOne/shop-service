package v1

import (
	"github.com/gorilla/mux"
	"shop-service/internal/service"
)

type Router struct {
	Mux *mux.Router
}

func New() *Router {
	return &Router{
		Mux: mux.NewRouter(),
	}
}

func NewRouter(router *Router, services *service.Services) {
	authMiddleware := &AuthMiddleware{services.Auth}

	newAuthRoutes(router, services.Auth)

	router.Mux.Use(authMiddleware.AuthMiddleware)
}

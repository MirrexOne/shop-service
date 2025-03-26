package v1

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"shop-service/internal/service"
	"strings"
)

const (
	userIdCtx = "userId"
)

type AuthMiddleware struct {
	authService service.Auth
}

func NewAuthMiddleware(authService service.Auth) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (h *AuthMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := bearerToken(r)
		if !ok {
			log.Errorf("AuthMiddleware: bearerToken: %v", ErrInvalidAuthHeader)
			newErrorResponse(http.StatusUnauthorized, ErrInvalidAuthHeader.Error())
			return
		}

		userId, err := h.authService.ParseToken(token)
		if err != nil {
			log.Errorf("AuthMiddleware: h.authService.ParseToken: %v", err)
			newErrorResponse(http.StatusUnauthorized, ErrCannotParseToken.Error())
			return
		}

		ctx := context.WithValue(r.Context(), userIdCtx, userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func bearerToken(r *http.Request) (string, bool) {
	const prefix = "Bearer "

	header := r.Header.Get("Authorization")
	if header == "" {
		return "", false
	}

	if len(header) > len(prefix) && strings.EqualFold(header[:len(prefix)], prefix) {
		return header[len(prefix):], true
	}

	return "", false
}

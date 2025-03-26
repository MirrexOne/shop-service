package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"shop-service/internal/service"
)

type authRoute struct {
	authService service.Auth
}

func newAuthRoutes(r *Router, authService service.Auth) {
	auth := &authRoute{
		authService: authService,
	}
	secureRouter := r.Mux.PathPrefix("/api/login").Subrouter()
	authMiddleware := NewAuthMiddleware(authService)
	secureRouter.Use(authMiddleware.AuthMiddleware)

	r.Mux.HandleFunc("/api/auth", auth.signUp).Methods("POST")
	secureRouter.HandleFunc("/api/login", auth.signIn).Methods("POST")
}

type signUpInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (auth *authRoute) signUp(w http.ResponseWriter, r *http.Request) {
	var input signUpInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}

	userId, err := auth.authService.CreateUser(r.Context(), service.AuthCreateUserInput{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			json.NewEncoder(w).Encode(newErrorResponse(http.StatusBadRequest, err.Error()))
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	token, err := auth.authService.GenerateToken(r.Context(), service.AuthGenerateTokenInput{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			http.Error(w, "invalid username or password", http.StatusBadRequest)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	type response struct {
		Code int `json:"code"`
		Id   int `json:"id"`
	}

	json.NewEncoder(w).Encode(response{
		Code: http.StatusCreated,
		Id:   userId,
	})
}

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (auth *authRoute) signIn(w http.ResponseWriter, r *http.Request) {
	var input signInInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

}

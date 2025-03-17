package v1

import (
	"encoding/json"
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

	r.Mux.HandleFunc("/auth/sign-up", auth.SignUp)
}

type signUpInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (auth *authRoute) SignUp(w http.ResponseWriter, r *http.Request) {
	var input signUpInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id, err := auth.authService.CreateUser(r.Context(), service.AuthCreateUserInput{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

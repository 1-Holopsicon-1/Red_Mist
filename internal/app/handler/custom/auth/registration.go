package handler

import (
	"RedMist/internal/app/dto"
	"RedMist/internal/app/services"
	"encoding/json"
	"net/http"
)

type Register struct {
}

func RegisterHandler(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input dto.RegisterDTO
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := userService.CreateUser(input.Username, input.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Генерируем токен
		authService := &AuthService{userRepository: repository.NewUserRepository()}
		token, err := authService.GenerateToken(user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Отправляем ответ с токеном
		response := TokenResponse{
			Token: token,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

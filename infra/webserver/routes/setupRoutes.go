package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaovds/chat/application/domain"
	"github.com/joaovds/chat/application/repository"
)

func SetupRoutes(mainMux *chi.Mux, userRepo *repository.UserRepository) {
	apiV1 := chi.NewRouter()

	handleWebsocketRoutes(apiV1)
	handleRoomRoutes(apiV1)
	handleUserRoutes(apiV1, userRepo)

	mainMux.Mount("/api/v1", apiV1)
}

func handleUserRoutes(mux *chi.Mux, userRepo *repository.UserRepository) {

	mux.Post("/createUser", func(writer http.ResponseWriter, request *http.Request) {
		var newUser domain.User
		if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
			http.Error(writer, "Error decodifying JSON", http.StatusBadRequest)
			return
		}

		if newUser.Gender == "" || newUser.Nickname == "" || newUser.Password == "" {
			http.Error(writer, "All mandatory fields must be provided.", http.StatusBadRequest)
			return
		}

		existingUser, err := userRepo.GetUserByNickName(newUser.Nickname)
		if existingUser != nil {
			http.Error(writer, "Nickname must be unique", http.StatusBadRequest)
			return
		} else if err != nil && err.Error() != "User not found" {
			http.Error(writer, "There was an error getting the user by the given nickname.", http.StatusInternalServerError)
			return
		}

		err = userRepo.CreateUser(&newUser)
		if err != nil {
			http.Error(writer, "Error creating new user", http.StatusInternalServerError)
			return
		}

		writer.Write([]byte("User created successfully."))
	})
}

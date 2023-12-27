package route

import (
	"github.com/BelyaevEI/GophKeeper/client/internal/data/dataservice"
	"github.com/BelyaevEI/GophKeeper/client/internal/middlewares"
	"github.com/BelyaevEI/GophKeeper/client/internal/user/userservice"
	"github.com/go-chi/chi"
)

func New(userservice *userservice.UserService,
	dataservice *dataservice.DataService,
	middlewares *middlewares.Middlewares) *chi.Mux {

	route := chi.NewRouter()

	// Handlers
	route.Post("/api/user/registration", userservice.Registration)   // Registration new user
	route.Post("/api/user/authorization", userservice.Authorization) // Authorization user

	// CRUD handlers
	route.Post("/api/user/authorization", middlewares.Authentication(dataservice.SaveData)) // Saving user data

	return route
}

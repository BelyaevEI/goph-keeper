package middlewares

import (
	"net/http"
	"strconv"

	"github.com/BelyaevEI/GophKeeper/server/internal/user/userservice"
)

type Middlewares struct {
	userService *userservice.UserService
}

func New(userService *userservice.UserService) *Middlewares {
	return &Middlewares{userService: userService}
}

func (middlewares *Middlewares) Authentication(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()

		// Read user information
		userIDs, exists := request.Header["Userid"]
		if !exists || len(userIDs) == 0 {
			middlewares.userService.Log.Log.Info("userID is empty")
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		userID, err := strconv.Atoi(userIDs[0])
		if err != nil {
			middlewares.userService.Log.Log.Error("the userID was determined with an error: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		token := request.Header["Token"][0]
		if len(token) == 0 {
			middlewares.userService.Log.Log.Error("user token is empty")
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validation user token
		ok, err := middlewares.userService.ValidationToken(ctx, uint32(userID), token)
		if err != nil {
			middlewares.userService.Log.Log.Error("validation user token with an error: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		// Verifying the validity of the token
		if !ok {
			middlewares.userService.Log.Log.Info("validation user token is failed")
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(writer, request.WithContext(ctx))
	})
}

package userservice

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/BelyaevEI/GophKeeper/client/internal/cookies"
	"github.com/BelyaevEI/GophKeeper/client/internal/logger"
	"github.com/BelyaevEI/GophKeeper/client/internal/models"
	"github.com/BelyaevEI/GophKeeper/client/internal/user/userrepository"
)

type UserService struct {
	Log            *logger.Logger
	userrepository *userrepository.UserRepository
}

func New(log *logger.Logger, userrepository *userrepository.UserRepository) *UserService {
	return &UserService{
		userrepository: userrepository,
		Log:            log,
	}
}

func (user *UserService) Registration(writer http.ResponseWriter, request *http.Request) {

	var (
		buf     bytes.Buffer
		regData models.RegistrationData
	)

	ctx := request.Context()

	// Given registration data from client for registration
	_, err := buf.ReadFrom(request.Body)
	if err != nil {
		user.Log.Log.Error("read body from request is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Deserializing JSON
	if err = json.Unmarshal(buf.Bytes(), &regData); err != nil {
		user.Log.Log.Error("deserializing JSON is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Checking unique login
	if unique := user.userrepository.CheckUniqueLogin(ctx, regData.Login); !unique {
		user.Log.Log.Info("this user has been registered")
		writer.WriteHeader(http.StatusConflict)
		return
	}

	// Creating unique userID for this user
	userID, err := user.userrepository.GenerateUniqueUserID()
	if err != nil {
		user.Log.Log.Error("creating unique userID is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Creating secret key for create token
	key := user.userrepository.GenerateRandomString(7)

	token, err := cookies.NewCookie(writer, userID, key)
	if err != nil {
		user.Log.Log.Error("creating token for cookie is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Saving login/password and secret key, new user
	err = user.userrepository.SaveDataNewUser(ctx, regData.Login, regData.Password, key)
	if err != nil {
		user.Log.Log.Error("saving new user into db is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send to client userID and token for authentication
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	response := models.RespRegistrationData{Token: token, UserID: userID}

	//Serializing response server
	enc := json.NewEncoder(writer)
	if err := enc.Encode(response); err != nil {
		user.Log.Log.Error("serializing response server is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (user *UserService) ValidationToken(ctx context.Context, userID uint32, token string) (bool, error) {

	// Get password for user validation
	secretKey, err := user.userrepository.GetSecretKey(ctx, userID)
	if err != nil {
		return false, err
	}

	// Check token is valid
	ok := cookies.Validation(token, secretKey)
	if !ok {
		return false, nil
	}
	return true, nil
}

func (user *UserService) Authorization(writer http.ResponseWriter, request *http.Request) {

	var (
		buf      bytes.Buffer
		authData models.RegistrationData
	)

	ctx := request.Context()

	// Given data from client for authorization
	_, err := buf.ReadFrom(request.Body)
	if err != nil {
		user.Log.Log.Error("read body from request is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Deserializing JSON
	if err = json.Unmarshal(buf.Bytes(), &authData); err != nil {
		user.Log.Log.Error("deserializing JSON is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verification of the user's password
	ok, err := user.userrepository.VerifyingPassword(ctx, authData.Login, authData.Password)
	if err != nil {
		user.Log.Log.Error("internal error with select from db: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if !ok {
		user.Log.Log.Info("verifying user is failed")
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Given userID for token
	userID, err := user.userrepository.GetUserID(ctx, authData.Login)
	if err != nil {
		user.Log.Log.Error("get userID is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Given secret key for token
	secretKey, err := user.userrepository.GetSecretKey(ctx, userID)
	if err != nil {
		user.Log.Log.Error("get secret key is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Creating token for user
	token, err := cookies.NewCookie(writer, userID, secretKey)
	if err != nil {
		user.Log.Log.Error("creating token for cookie is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send to client userID and token for authentication
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	response := models.RespRegistrationData{Token: token, UserID: userID}

	//Serializing response server
	enc := json.NewEncoder(writer)
	if err := enc.Encode(response); err != nil {
		user.Log.Log.Error("serializing response server is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

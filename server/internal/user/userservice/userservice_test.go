package userservice

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/BelyaevEI/GophKeeper/server/internal/logger"
	"github.com/BelyaevEI/GophKeeper/server/internal/models"
	dbmocks "github.com/BelyaevEI/GophKeeper/server/internal/storage/userdb/mocks"
	"github.com/BelyaevEI/GophKeeper/server/internal/user/userrepository"
	usermocks "github.com/BelyaevEI/GophKeeper/server/internal/user/userrepository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestRegistration(t *testing.T) {
	testcase := struct {
		name     string
		wantCode int
	}{
		name:     "Registration new user",
		wantCode: http.StatusOK,
	}

	t.Run(testcase.name, func(t *testing.T) {

		ctx := context.Background()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := dbmocks.NewMockStore(ctrl)
		user := usermocks.NewMockStore(ctrl)

		log, _ := logger.New()
		userrepository := userrepository.New(db)
		userservice := New(log, userrepository)

		user.EXPECT().CheckUniqueLogin(ctx, "Test").Return(true)
		db.EXPECT().CheckUniqueLogin(ctx, "Test").Return(nil)
		user.EXPECT().GenerateUniqueUserID().Return(uint32(1), nil)
		user.EXPECT().GenerateRandomString(7).Return("1234567")
		user.EXPECT().SaveDataNewUser(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		db.EXPECT().SaveDataNewUser(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		writer := httptest.NewRecorder()

		r := models.RegistrationData{
			Login:    "Test",
			Password: "Test"}

		req, _ := json.Marshal(r)
		request := httptest.NewRequest(http.MethodPost, "/api/user/registration", strings.NewReader(string(req)))
		request.Header.Set("Content-Type", "json/application")

		userservice.Registration(writer, request)

		result := writer.Result()
		defer result.Body.Close()

		assert.Equal(t, result.StatusCode, testcase.wantCode)

	})
}

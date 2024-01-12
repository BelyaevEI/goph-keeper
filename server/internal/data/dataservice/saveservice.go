package dataservice

import (
	"io"
	"net/http"

	"github.com/BelyaevEI/GophKeeper/server/internal/data/datarepository/bankrepository"
	"github.com/BelyaevEI/GophKeeper/server/internal/data/datarepository/binrepository"
	"github.com/BelyaevEI/GophKeeper/server/internal/data/datarepository/passwordrepository"
	"github.com/BelyaevEI/GophKeeper/server/internal/data/datarepository/textrepository"
	"github.com/BelyaevEI/GophKeeper/server/internal/logger"
)

type DataService struct {
	log                *logger.Logger
	Passwordrepository *passwordrepository.PasswordRepository
	Textrepository     *textrepository.TextRepository
	Binrepository      *binrepository.BinRepository
	Bankrepository     *bankrepository.BankRepository
}

func New(log *logger.Logger,
	passwordrepository *passwordrepository.PasswordRepository,
	textrepository *textrepository.TextRepository,
	binrepositroy *binrepository.BinRepository,
	bankrepository *bankrepository.BankRepository,
) *DataService {

	return &DataService{
		Passwordrepository: passwordrepository,
		Textrepository:     textrepository,
		Binrepository:      binrepositroy,
		Bankrepository:     bankrepository,
		log:                log,
	}
}

// Saving input user data
func (dataservice *DataService) SaveData(writer http.ResponseWriter, request *http.Request) {

	ctx := request.Context()

	// Getting data to save
	body, err := io.ReadAll(request.Body)
	if err != nil {
		dataservice.log.Log.Error("reading body from request is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Getting headers to define the data format
	// Saving user's login/password of service
	logins := request.Header["Login"]
	if len(logins) > 0 {
		err := dataservice.Passwordrepository.SaveData(ctx, body)
		if err != nil {
			dataservice.log.Log.Error("saving login/password is failed: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
		return
	}

	// Saving user's text of service
	texts := request.Header["Text"]
	if len(texts) > 0 {
		err := dataservice.Textrepository.SaveData(ctx, body)
		if err != nil {
			dataservice.log.Log.Error("saving text is failed: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
		return
	}

	// Saving user's binary data of service
	bins := request.Header["Bin"]
	if len(bins) > 0 {
		err := dataservice.Binrepository.SaveData(ctx, body)
		if err != nil {
			dataservice.log.Log.Error("saving binary data is failed: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
		return
	}

	// Saving user's bank data of service
	banks := request.Header["Bank"]
	if len(banks) > 0 {
		err := dataservice.Bankrepository.SaveData(ctx, body)
		if err != nil {
			dataservice.log.Log.Error("saving binary data is failed: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
		return
	}

	dataservice.log.Log.Info("empty client header and not define type of data")
	writer.WriteHeader(http.StatusBadRequest)
}

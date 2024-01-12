package dataservice

import (
	"io"
	"net/http"
)

// Reading user data
func (dataservice *DataService) UpdateData(writer http.ResponseWriter, request *http.Request) {

	ctx := request.Context()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		dataservice.log.Log.Error("reading body from request is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Getting headers to define the data format
	// Updating user's login/passwords of service
	logins := request.Header["Login"]
	if len(logins) > 0 {
		if err := dataservice.Passwordrepository.UpdateData(ctx, body); err != nil {
			dataservice.log.Log.Error("updating data is failed: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}

	// Updating user's text of service
	texts := request.Header["Text"]
	if len(texts) > 0 {
		if err := dataservice.Textrepository.UpdateData(ctx, body); err != nil {
			dataservice.log.Log.Error("updating data is failed: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}

	// Updating user's bin data of service
	bins := request.Header["Bin"]
	if len(bins) > 0 {
		if err := dataservice.Binrepository.UpdateData(ctx, body); err != nil {
			dataservice.log.Log.Error("updating data is failed: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}

	// Saving user's bank data of service
	banks := request.Header["Bank"]
	if len(banks) > 0 {
		if err := dataservice.Bankrepository.UpdateData(ctx, body); err != nil {
			dataservice.log.Log.Error("updating data is failed: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}

	dataservice.log.Log.Info("empty client header and not define type of data")
	writer.WriteHeader(http.StatusBadRequest)
}

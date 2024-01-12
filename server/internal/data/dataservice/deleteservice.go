package dataservice

import (
	"io"
	"net/http"
)

// Deleting user data
func (dataservice *DataService) DeleteData(writer http.ResponseWriter, request *http.Request) {

	ctx := request.Context()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		dataservice.log.Log.Error("reading body from request is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Getting headers to define the data format
	// Deleting user's login/passwords of service
	logins := request.Header["Login"]
	if len(logins) > 0 {
		if err := dataservice.Passwordrepository.DeleteData(ctx, body); err != nil {
			dataservice.log.Log.Error("updating data is failed: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}

	// Deleting user's text of service
	texts := request.Header["Text"]
	if len(texts) > 0 {
		if err := dataservice.Textrepository.DeleteData(ctx, body); err != nil {
			dataservice.log.Log.Error("updating data is failed: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}

	// Deleting user's bin data of service
	bins := request.Header["Bin"]
	if len(bins) > 0 {
		if err := dataservice.Binrepository.DeleteData(ctx, body); err != nil {
			dataservice.log.Log.Error("updating data is failed: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}

	// Deleting user's bank data of service
	banks := request.Header["Bank"]
	if len(banks) > 0 {
		if err := dataservice.Bankrepository.DeleteData(ctx, body); err != nil {
			dataservice.log.Log.Error("updating data is failed: ", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}

	dataservice.log.Log.Info("empty client header and not define type of data")
	writer.WriteHeader(http.StatusBadRequest)
}

package dataservice

import (
	"database/sql"
	"errors"
	"io"
	"net/http"
)

// Reading user data
func (dataservice *DataService) ReadData(writer http.ResponseWriter, request *http.Request) {

	ctx := request.Context()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		dataservice.log.Log.Error("reading body from request is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Getting headers to define the data format
	// Returning user's login/passwords of service
	logins := request.Header["Login"]
	if len(logins) > 0 {
		data, err := dataservice.Passwordrepository.GetData(ctx, body)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				dataservice.log.Log.Error("reading login/passwords data is failed: ", err)
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			dataservice.log.Log.Error("data not found: ", err)
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write(data)
		return
	}

	// Returning user's text of service
	texts := request.Header["Text"]
	if len(texts) > 0 {
		data, err := dataservice.Textrepository.GetData(ctx, body)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				dataservice.log.Log.Error("reading texts data is failed: ", err)
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			dataservice.log.Log.Error("data not found: ", err)
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write(data)
		return
	}

	// Returning user's bin data of service
	bins := request.Header["Bin"]
	if len(bins) > 0 {

		data, err := dataservice.Binrepository.GetData(ctx, body)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				dataservice.log.Log.Error("reading binary data is failed: ", err)
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			dataservice.log.Log.Error("data not found: ", err)
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write(data)
		return

	}

	// Saving user's bank data of service
	banks := request.Header["Bank"]
	if len(banks) > 0 {
		data, err := dataservice.Bankrepository.GetData(ctx, body)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				dataservice.log.Log.Error("reading bank data is failed: ", err)
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			dataservice.log.Log.Error("data not found: ", err)
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write(data)
		return
	}

	dataservice.log.Log.Info("empty client header and not define type of data")
	writer.WriteHeader(http.StatusBadRequest)
}

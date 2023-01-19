package server

import (
	"encoding/json"
	"github.com/margostino/earth-station-api/common"
	"net/http"
)

func WriteResponse(payload interface{}, writer http.ResponseWriter) {
	response, err := json.Marshal(payload)
	if !common.IsError(err, "error happened in JSON marshal") {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(response)
	}
}

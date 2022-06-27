package handler

import (
	"encoding/json"
	"gokiosk/internal/errors"
	"log"
	"net/http"
	"strconv"
)

// GetPagination will read offset and limit from the request
func getPaginationParams(r *http.Request) (int, int) {
	offset := 0
	limit := 10
	if r.URL.Query().Get("offset") != "" {
		offset, _ = strconv.Atoi(r.URL.Query().Get("offset"))
	}
	if r.URL.Query().Get("limit") != "" {
		limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	}
	return offset, limit
}

func writeJsonResponse(w http.ResponseWriter, status int, data interface{}) {
	result, err := json.Marshal(data)
	if err != nil {
		log.Printf("ERROR MARSHALING JSON: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(result)
}

func writeErrorResponse(w http.ResponseWriter, status int, err error) {
	log.Printf("ERROR: %s", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}

func writeInternalServerError(w http.ResponseWriter) {
	appErr := errors.NewAppError("Internal Server Error", "internal_server_error")
	writeErrorResponse(w, http.StatusInternalServerError, appErr)
}

func writeBadRequestError(w http.ResponseWriter, message, code string) {
	appErr := errors.NewAppError(message, code)
	writeErrorResponse(w, http.StatusBadRequest, appErr)
}

package handler

import (
	"encoding/json"
	"fmt"
	"gokiosk/internal/errors"
	"gokiosk/internal/model"
	"log"
	"net/http"
	"strconv"
)

// GetPagination will read offset and limit from the request
func getPaginationParams(r *http.Request) model.Paginate {
	offset := 0
	limit := 10
	if r.URL.Query().Get("offset") != "" {
		offset, _ = strconv.Atoi(r.URL.Query().Get("offset"))
	}
	if r.URL.Query().Get("limit") != "" {
		limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	}
	return model.Paginate{Limit: offset, Offset: limit}
}

func writeJsonResponse(w http.ResponseWriter, status int, data interface{}) {
	result, err := json.Marshal(data)
	if err != nil {
		log.Printf("ERROR MARSHALING JSON: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if data != nil {
		w.Write(result)
	}
}

func writeErrorResponse(w http.ResponseWriter, status int, err error) {
	log.Printf("ERROR: %s", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}

func writeInternalServerError(w http.ResponseWriter) {
	appErr := fmt.Errorf(errors.ERR_INTERNAL_SERVER_ERROR)
	writeErrorResponse(w, http.StatusInternalServerError, appErr)
}

func writeBadRequestError(w http.ResponseWriter) {
	appErr := fmt.Errorf(errors.ERR_BAD_REQUEST)
	writeErrorResponse(w, http.StatusBadRequest, appErr)
}

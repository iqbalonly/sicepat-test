package api

import (
	"bytes"
	"encoding/json"
	"errors"

	"io"
	"log"
	"net/http"
	"sicepat/internal/constant"
	"sicepat/internal/dto"
	"sicepat/internal/ierr"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	headerKeyContentType = "Content-Type"
	headerKeyTS          = "X-Timestamp"
)

// renderResponse defines a wrapper for sending JSON response.
func renderResponse(w http.ResponseWriter, resp interface{}, httpCode int, iCode constant.Code) {
	w.Header().Set(headerKeyContentType, "application/json")
	w.Header().Set(headerKeyTS, time.Now().Format(time.RFC3339))
	w.WriteHeader(httpCode)

	if err := json.NewEncoder(w).Encode(&dto.BaseResponse{
		Code: int(iCode),
		Msg:  "Success",
		Data: resp,
	}); err != nil {
		panic(err)
	}
}

// renderErrorResponse defines a wrapper for sending error JSON response.
func renderErrorResponse(w http.ResponseWriter, err error, httpResponseCode int) {
	w.Header().Set(headerKeyContentType, "application/json")
	w.Header().Set(headerKeyTS, time.Now().Format(time.RFC3339))

	var iErr *ierr.Error
	if errors.As(err, &iErr) {
		w.WriteHeader(iErr.GetHTTPCode())

		if rErr := json.NewEncoder(w).Encode(&dto.BaseResponse{
			Code: int(iErr.GetCode()),
			Msg:  iErr.Error(),
		}); rErr != nil {
			panic(err)
		}

		return
	}

	w.WriteHeader(httpResponseCode)
	if rErr := json.NewEncoder(w).Encode(dto.BaseResponse{
		Code: httpResponseCode,
		Msg:  err.Error(),
	}); rErr != nil {
		panic(err)
	}
}

// parseAPIRequest parses and validate the incoming HTTP request.
func parseAPIRequest(r *http.Request, req dto.RequestAble) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		panic(err)
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err = json.NewDecoder(io.NopCloser(bytes.NewBuffer(body))).Decode(req)
	if err != nil {
		logrus.Error("failed to decode request")
		return err
	}

	err = req.Validate()
	if err != nil {
		return err
	}

	return nil
}

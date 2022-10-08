package http

import (
	"context"
	"encoding/json"
	"errors"
	"go-kit-reddit-demo/internal/reddit/service"
	http1 "net/http"
)

func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

func err2code(err error) int {
	switch err {
	case service.ErrForbidden:
		return http1.StatusForbidden
	}
	return http1.StatusBadRequest
}

type errorWrapper struct {
	Error string `json:"error"`
}

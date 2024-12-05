package main

import (
	"log/slog"
	"net/http"
)

func(a *App) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri = r.URL.RequestURI()
	)

	a.logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
package handler

import (
	"net/http"
	"text/template"
)

// HomeHandler handles the _healthcheck_ endpoint that is the standard for health check
type HomeHandler struct {
}

// NewHomeHandler return a instance of a HomeHandler
func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

// ServeHTTP handle the http request, it prints 200 and return manually the current system version
func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

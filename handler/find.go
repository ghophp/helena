package handler

import (
	"net/http"
	"text/template"
)

type FindHandler struct {
}

// NewFindHandler return a instance of a FindHandler
func NewFindHandler() *FindHandler {
	return &FindHandler{}
}

// ServeHTTP handle the http request, it prints 200 and return manually the current system version
func (h *FindHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

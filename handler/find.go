package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ghophp/helena/db/playlist"
	"github.com/ghophp/helena/db/track"

	"gopkg.in/gorp.v1"
)

type (
	FindHandler struct {
		db gorp.SqlExecutor
	}

	ClientProperties struct {
		Genres   []string `json:"genres"`
		Emotions []string `json:"emotions"`
	}
)

// NewFindHandler return a instance of a FindHandler
func NewFindHandler(db gorp.SqlExecutor) *FindHandler {
	return &FindHandler{db}
}

// ServeHTTP handle the http request, it prints 200 and return manually the current system version
func (h *FindHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var props ClientProperties
	err = json.Unmarshal(body, &props)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if len(props.Emotions) <= 0 || len(props.Genres) <= 0 {
		http.Error(w, "invalid number of emotions or genres", http.StatusBadRequest)
	}

	playlists, err := playlist.GetPlaylistByGenres(h.db, props.Genres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if len(playlists) <= 0 {
		playlists, err = playlist.GetAllPlaylists(h.db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	tracks, err := track.GetAllTracksByPlaylists(playlists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//for _, t := range tracks {
	// todo: implement the filter logic that with the article studies
	//}

}

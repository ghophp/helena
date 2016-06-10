package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

	var playlistIds []string
	for _, p := range playlists {
		playlistIds = append(playlistIds, strconv.FormatUint(p.Id, 10))
	}

	tracks, err := track.GetAllTracksByPlaylists(h.db, playlistIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	result, err := json.Marshal(tracks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(result))
}

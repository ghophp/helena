package track

import (
	"database/sql"
	"strconv"

	"gopkg.in/gorp.v1"
)

// Track is the struct that maps the table track
type Track struct {
	Id               uint64        `db:"id"`
	PlaylistId       uint64        `db:"playlist_id"`
	Href             string        `db:"href"`
	Duration         uint64        `db:"duration"`
	Popularity       uint64        `db:"popularity"`
	Reference        string        `db:"reference"`
	Danceability     float64       `db:"danceability"`
	Energy           float64       `db:"energy"`
	Loudness         float64       `db:"loudness"`
	Speechiness      float64       `db:"speechiness"`
	Mode             uint64        `db:"mode"`
	Key              uint64        `db:"track_key"`
	Acousticness     float64       `db:"acousticness"`
	Instrumentalness float64       `db:"instrumentalness"`
	Liveness         float64       `db:"liveness"`
	Valence          float64       `db:"valence"`
	Tempo            float64       `db:"tempo"`
	TimeSignature    uint64        `db:"time_signature"`
	LastUpdate       sql.NullInt64 `db:"last_update"`
}

func GetAllTracksByPlaylists(db gorp.SqlExecutor, playlistId uint64) ([]*Playlist, error) {
	var (
		query  = "SELECT * FROM track WHERE playlist_id = " + strconv.FormatUint(playlistId, 10)
		tracks []*Track
	)

	_, err := db.Select(&tracks, query)
	return tracks, err
}

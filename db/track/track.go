package track

import (
	"database/sql"
	"strings"

	"gopkg.in/gorp.v1"
)

// Track is the struct that maps the table track
type Track struct {
	Id               uint64        `json:"id" db:"id"`
	PlaylistId       uint64        `json:"playlist_id" db:"playlist_id"`
	EmotionId        uint64        `json:"-" db:"emotion_id"`
	Href             string        `json:"href" db:"href"`
	Duration         uint64        `json:"-" db:"duration"`
	Popularity       uint64        `json:"-" db:"popularity"`
	Reference        string        `json:"reference" db:"reference"`
	Danceability     float64       `json:"-" db:"danceability"`
	Energy           float64       `json:"-" db:"energy"`
	Loudness         float64       `json:"-" db:"loudness"`
	Speechiness      float64       `json:"-" db:"speechiness"`
	Mode             uint64        `json:"-" db:"mode"`
	Key              uint64        `json:"-" db:"track_key"`
	Acousticness     float64       `json:"-" db:"acousticness"`
	Instrumentalness float64       `json:"-" db:"instrumentalness"`
	Liveness         float64       `json:"-" db:"liveness"`
	Valence          float64       `json:"-" db:"valence"`
	Tempo            float64       `json:"-" db:"tempo"`
	TimeSignature    uint64        `json:"-" db:"time_signature"`
	LastUpdate       sql.NullInt64 `json:"-" db:"last_update"`
}

func GetAllTracksByPlaylists(db gorp.SqlExecutor, playlistIds []string) ([]*Track, error) {
	var (
		query  = "SELECT * FROM track WHERE playlist_id IN (" + strings.Join(playlistIds, ",") + ")"
		tracks []*Track
	)

	_, err := db.Select(&tracks, query)
	return tracks, err
}

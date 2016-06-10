package playlist

import (
	"database/sql"
	"strings"

	"gopkg.in/gorp.v1"
)

// Playlist is the struct that maps the table playlist
type Playlist struct {
	Id            uint64        `db:"id"`
	Name          string        `db:"name"`
	Reference     string        `db:"reference"`
	UserReference string        `db:"user_reference"`
	LastUpdate    sql.NullInt64 `db:"last_update"`
}

func GetAllPlaylists(db gorp.SqlExecutor) ([]*Playlist, error) {
	var (
		query     = "SELECT * FROM playlist WHERE last_update IS NOT NULL"
		playlists []*Playlist
	)

	_, err := db.Select(&playlists, query)
	return playlists, err
}

func GetPlaylistByGenres(db gorp.SqlExecutor, genres []string) ([]*Playlist, error) {
	for i, _ := range genres {
		genres[i] = `"` + genres[i] + `"`
	}

	var (
		query = `SELECT p.* FROM playlist AS p 
			INNER JOIN playlist_genre AS pg ON pg.playlist_id = p.id 
			INNER JOIN genre AS g ON g.id = pg.genre_id 
			WHERE g.name IN (` + strings.Join(genres, ",") + `) AND p.last_update IS NOT NULL`

		playlists []*Playlist
	)

	_, err := db.Select(&playlists, query)
	return playlists, err
}

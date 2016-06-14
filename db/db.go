// Package db handles the database management and mapping, through the
// structs on each subpackage and management funcs we interact with the
// database and perform the queries
package db

import (
	"database/sql"

	"github.com/ghophp/helena/db/playlist"
	"github.com/ghophp/helena/db/track"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

// NewDb return a new instance of gorp.DbMap or an error
// if was not possible to open the connection or there
// was some problem mapping
func NewDb(connection string) (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbMap.AddTableWithName(playlist.Playlist{}, "playlist").SetKeys(true, "id")
	dbMap.AddTableWithName(track.Track{}, "track").SetKeys(true, "id")

	return dbMap, nil
}

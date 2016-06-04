// Package config handles the parsing and marshall of the configuration file
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

const (
	// ConfigDirEnv is the name of the env variable that contains the path for the config.json file
	ConfigDirEnv = "HELENA_CONFIG_DIR"
)

type (
	// Config is the struct used to unmarshal the config.json file
	Config struct {
		Port     int       `json:"port"`
		Database *database `json:"database"`
	}

	database struct {
		Name     string `json:"name"`
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
	}
)

// NewConfig find the config.json file, unmarshal it and return a new instance of Config
func NewConfig() (*Config, error) {
	dir := os.Getenv(ConfigDirEnv)
	if len(dir) <= 0 {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dir = path.Join(pwd, "/config.json")
	}

	bytes, err := ioutil.ReadFile(dir)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = json.Unmarshal(bytes, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// String return the database connection string eg. user:password@tcp(localhost:5555)/dbname
func (db *database) String() string {
	return db.User + ":" +
		db.Password + "@tcp(" +
		db.Host + ":" +
		strconv.Itoa(db.Port) + ")/" +
		db.Name
}

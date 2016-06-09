package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ghophp/helena/config"
	"github.com/ghophp/helena/db"
	"github.com/ghophp/helena/handler"
	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbMap, err := db.NewDb(cfg.Database.String())
	if err != nil {
		log.Fatal(err)
	}

	homeHandler := handler.NewHomeHandler()
	findHandler := handler.NewFindHandler(dbMap)

	r := mux.NewRouter()
	r.Handle("/", homeHandler)
	r.Handle("/find", findHandler)

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./dist/")))
	r.PathPrefix("/static/").Handler(s)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r))
}

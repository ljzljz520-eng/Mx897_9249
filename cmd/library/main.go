package main

import (
	"librarytransfer/api"
	"librarytransfer/config"
	"librarytransfer/storage"
	"log"
	"net/http"
)

func main() {
	c := config.Load()
	if e := config.Validate(c); e != nil {
		log.Fatal(e)
	}
	s, e := storage.Open(c.DataFile)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	log.Printf("library reader transfer listening on %s", c.ListenAddr)
	log.Fatal(http.ListenAndServe(c.ListenAddr, api.Handler{Store: s}))
}

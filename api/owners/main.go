package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/owners", handlerFunc)

	server := http.Server{
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	log.Fatal(server.ListenAndServe())
}

func handlerFunc(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	default:
		{
			res.WriteHeader(501)
			res.Write([]byte("not implemented"))
		}
	}
}

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := fetchNameAndJoke()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(res))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	log.Fatal(server.ListenAndServe())
}

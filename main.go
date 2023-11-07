package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fetch our name and joke with 5 second limit before timeout
		res, err := fetchNameAndJoke(5 * time.Second)
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

package main

import (
	"./movie"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	movieRepository := movie.Repository{}

	data := movieRepository.GetByPage(1)

	movies, _ := json.Marshal(data)

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(movies)
	})

	http.Handle("/", r)

	http.ListenAndServe(":3000", nil)

}

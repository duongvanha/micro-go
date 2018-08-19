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

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data, err := movieRepository.GetByPage(1)

		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

		movies, err := json.Marshal(data)

		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

		writer.Write(movies)
	})

	http.Handle("/", r)

	http.ListenAndServe(":3000", nil)

}

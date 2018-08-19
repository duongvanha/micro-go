package main

import (
	"./movie"
	"encoding/json"
	"fmt"
)

func main() {

	movieRepository := movie.Repository{}

	data, _ := json.Marshal(movieRepository.GetByPage(1))

	fmt.Println(data)
}

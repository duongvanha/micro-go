package main

import (
	"./auth"
	"fmt"
)

func main() {
	user := auth.User{Id: 12, UserName: "name", Password: "name"}
	authService := auth.Service{}

	authService.Register("sieunhangao	", "123456")

	authService.Login("sieunhangao", "123456")

	fmt.Println(user.ToString())
}

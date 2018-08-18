package main

import (
	"./authservice"
	"fmt"
)

func main() {
	user := authservice.User{Id: 12, UserName: "name", Password: "name"}
	auth := authservice.AuthService{}

	auth.Register("sieunhangao	", "123456")

	auth.Login("sieunhangao", "123456")

	fmt.Println(user.ToString())
}

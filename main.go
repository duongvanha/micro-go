package main

import (
	"fmt"
	"github.com/duongvanha/micro-go/authService"
)

func main() {
	user := authservice.User{Id: 12, UserName: "name", Password: "name"}

	fmt.Println(user.ToString())
}

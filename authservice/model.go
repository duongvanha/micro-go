package authservice

import "fmt"

type User struct {
	id       int
	userName string
	password string
}

func (u *User) toString() string {
	return fmt.Sprintf("id : %d, userName: %s", u.id, u.userName)
}

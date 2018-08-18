package authService

import "errors"

type User struct {
	userName string
	password string
}

func (user User) equal(userName string, password string) bool {
	return user.userName == userName && user.password == password
}

type AuthService struct {
	users []User
}

func (auth *AuthService) login(userName string, password string) (User, error) {

	for _, user := range auth.users {
		if user.equal(userName, password) {
			return user, nil
		}
	}

	return User{}, errors.New(userName + " not found")
}

func (auth *AuthService) register(userName string, password string) error {
	for _, user := range auth.users {
		if user.equal(userName, password) {
			return errors.New("the username is already is used")
		}
	}

	return nil
}

package auth

import "errors"

type Service struct {
	users []User
}

func (auth *Service) Login(userName string, password string) (User, error) {

	for _, user := range auth.users {
		if user.UserName == userName && user.Password == password {
			return user, nil
		}
	}

	return User{}, errors.New(userName + " not found")
}

func (auth *Service) Register(userName string, password string) error {
	for _, user := range auth.users {
		if user.UserName == userName && user.Password == password {
			return errors.New("the username is already is used")
		}
	}

	return nil
}

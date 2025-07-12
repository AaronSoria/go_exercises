package mystructs

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreateAt  time.Time
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("fields can't be empty")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreateAt:  time.Now(),
	}, nil
}

func (user User) ShowUserInfo() {
	fmt.Println(user.FirstName)
	fmt.Println(user.LastName)
	fmt.Println(user.BirthDate)
	fmt.Println(user.CreateAt)
}

func (user *User) Clear() {
	user.FirstName = ""
	user.LastName = ""
	user.BirthDate = ""
}

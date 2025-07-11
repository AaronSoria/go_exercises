package mystructs

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreateAt  time.Time
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

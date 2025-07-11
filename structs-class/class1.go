package main

import (
	"fmt"
	"time"

	"example.com/structs-class/mystructs"
)

func main() {
	firstName := getStringInput("Enter first name")
	lastName := getStringInput("Enter last name")
	birthDate := getStringInput("Enter birthDate")
	var user mystructs.User
	user = mystructs.User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreateAt:  time.Now(),
	}

	user.ShowUserInfo()
	user.Clear()
	user.ShowUserInfo()
}

func getStringInput(message string) (output string) {
	fmt.Println(message)
	fmt.Scan(&output)
	return output
}

func showUserPointerInfo(user *mystructs.User) {
	user.FirstName = "Luis Aaron"
	fmt.Println(user.FirstName)
	fmt.Println(user.LastName)
	fmt.Println(user.BirthDate)
	fmt.Println(user.CreateAt)
}

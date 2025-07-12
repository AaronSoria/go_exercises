package main

import (
	"fmt"

	"example.com/structs-class/mystructs"
)

func main() {
	firstName := getStringInput("Enter first name")
	lastName := getStringInput("Enter last name")
	birthDate := getStringInput("Enter birthDate")
	user, err := mystructs.New(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	user.ShowUserInfo()
	user.Clear()
	user.ShowUserInfo()

	admin := mystructs.NewAdmin("admin@boss.com", "1234", firstName, lastName, birthDate)
	admin.ShowUserInfo()
}

func getStringInput(message string) (output string) {
	fmt.Println(message)
	fmt.Scanln(&output)
	return output
}

func showUserPointerInfo(user *mystructs.User) {
	user.FirstName = "Luis Aaron"
	fmt.Println(user.FirstName)
	fmt.Println(user.LastName)
	fmt.Println(user.BirthDate)
	fmt.Println(user.CreateAt)
}

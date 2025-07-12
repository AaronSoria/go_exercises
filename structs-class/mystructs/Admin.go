package mystructs

import "time"

type Admin struct {
	Email    string
	Password string
	//User     User
	User
}

func NewAdmin(email, password, firstName, lastName, birthDate string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: firstName,
			LastName:  lastName,
			CreateAt:  time.Now(),
			BirthDate: birthDate,
		},
	}
}

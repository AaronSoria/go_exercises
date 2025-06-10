package main

import "fmt"

func main() {

	var accountBalance float64 = 1000.0

	fmt.Println("Welcome to Go bank!")
	fmt.Println("What do yoy want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance: ", accountBalance)
	} else if choice == 2 {
		var deposit float64
		print("Your deposit: ")
		fmt.Scan(&deposit)
		accountBalance += deposit
		fmt.Println("Balance updated! New amount: ", accountBalance)
	} else if choice == 3 {
		var Withdraw float64
		print("Your Withdraw: ")
		fmt.Scan(&Withdraw)
		accountBalance -= Withdraw
		fmt.Println("Balance updated! New amount: ", accountBalance)
	}

	fmt.Println("Your choice:", choice)
}

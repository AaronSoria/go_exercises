package main

import (
	"fmt"

	"example.com/bank/fileOps"
)

const accountBalanceFile = "balance.txt"

func main() {
	// var accountBalance float64 = 1000.0
	accountBalance, err := fileOps.GetFloatFromFile(accountBalanceFile, 0)
	if err != nil {
		fmt.Println(err)
		panic("can't continue")
	}
	fileOps.WriteValueToFile(atmWithSwitch(accountBalance), accountBalanceFile)
}

func atmWithSwitch(accountBalance float64) float64 {
	exit := false
	for !exit {
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your balance: ", accountBalance)
		case 2:
			var deposit float64
			print("Your deposit: ")
			fmt.Scan(&deposit)
			if deposit < 0 {
				fmt.Print("invalid amount. Must be greater than 0")
				continue
			}
			accountBalance += deposit
			fmt.Println("Balance updated! New amount: ", accountBalance)
		case 3:
			var Withdraw float64
			print("Your Withdraw: ")
			fmt.Scan(&Withdraw)
			if Withdraw < 0 {
				fmt.Print("invalid amount. Must be greater than 0")
				continue
			}

			if Withdraw > accountBalance {
				fmt.Print("invalid amount. You can't widthdraw than you have")
				continue
			}
			accountBalance -= Withdraw
			fmt.Println("Balance updated! New amount: ", accountBalance)
		default:
			fmt.Println("Good bye!")
			exit = true
		}
	}

	fmt.Println("thanks for use our bank!")
	return accountBalance
}

// func atmWithIf(accountBalance float64) {

// 	for {
// 		fmt.Println("Welcome to Go bank!")
// 		fmt.Println("What do yoy want to do?")
// 		fmt.Println("1. Check balance")
// 		fmt.Println("2. Deposit money")
// 		fmt.Println("3. Withdraw money")
// 		fmt.Println("4. Exit")

// 		var choice int
// 		fmt.Print("Your choice: ")
// 		fmt.Scan(&choice)

// 		if choice == 1 {
// 			fmt.Println("Your balance: ", accountBalance)
// 		} else if choice == 2 {
// 			var deposit float64
// 			print("Your deposit: ")
// 			fmt.Scan(&deposit)
// 			if deposit < 0 {
// 				fmt.Print("invalid amount. Must be greater than 0")
// 				continue
// 			}
// 			accountBalance += deposit
// 			fmt.Println("Balance updated! New amount: ", accountBalance)
// 		} else if choice == 3 {
// 			var Withdraw float64
// 			print("Your Withdraw: ")
// 			fmt.Scan(&Withdraw)
// 			if Withdraw < 0 {
// 				fmt.Print("invalid amount. Must be greater than 0")
// 				continue
// 			}

// 			if Withdraw > accountBalance {
// 				fmt.Print("invalid amount. You can't widthdraw than you have")
// 				continue
// 			}
// 			accountBalance -= Withdraw
// 			fmt.Println("Balance updated! New amount: ", accountBalance)
// 		} else {
// 			fmt.Println("Good bye!")
// 			//return

// 			break
// 		}
// 	}

// 	fmt.Println("thanks for use our bank!")
// }

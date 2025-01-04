package main

import (
	"fmt"
	"project-bank/bank"
)

func main() {
	// ConsolePrinter 사용
	consolePrinter := &bank.ConsolePrinter{}
	account := bank.NewAccount(0.0, consolePrinter)

	// 또는 FilePrinter 사용
	// filePrinter, err := bank.NewFilePrinter("transactions.log")
	// if err != nil {
	// 	fmt.Println("Error creating file printer:", err)
	// 	return
	// }
	// defer filePrinter.Close()
	// account := bank.NewAccount(0.0, filePrinter)

	var choice int
	var amount float64

	for {
		fmt.Println("\nChoose an option: 1 (Deposit), 2 (Withdraw), 3 (Check Balance), 4 (Exit):")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 2:
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 3:
			account.Balance()
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			account.Error("Invalid option. Please choose again.")
		}
	}
}

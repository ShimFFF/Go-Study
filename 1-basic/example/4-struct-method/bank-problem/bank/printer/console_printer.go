package printer

import "fmt"

// ConsolePrinter: 콘솔 출력 구현체
type ConsolePrinter struct{}

func (c *ConsolePrinter) PrintDeposit(amount, balance float64) {
	fmt.Printf("Deposited: %.2f | New Balance: %.2f\n", amount, balance)
}

func (c *ConsolePrinter) PrintWithdraw(amount, balance float64) {
	fmt.Printf("Withdrew: %.2f | New Balance: %.2f\n", amount, balance)
}

func (c *ConsolePrinter) PrintError(message string) {
	fmt.Println("Error:", message)
}

func (c *ConsolePrinter) PrintBalance(balance float64) {
	fmt.Printf("Current Balance: %.2f\n", balance)
}

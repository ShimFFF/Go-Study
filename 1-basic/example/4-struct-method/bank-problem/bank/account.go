package bank

import "project-bank/bank/printer"

// Account 구조체 정의
type Account struct {
	balance float64
	printer printer.AccountOutput // 출력 전략 인터페이스
}

// NewAccount: Account 생성자
func NewAccount(initialBalance float64, printer printer.AccountOutput) *Account {
	return &Account{
		balance: initialBalance,
		printer: printer,
	}
}

// Deposit 메서드
func (a *Account) Deposit(amount float64) {
	if amount <= 0 {
		a.printer.PrintError("Invalid deposit amount. Must be greater than zero.")
		return
	}
	a.balance += amount
	a.printer.PrintDeposit(amount, a.balance)
}

// Withdraw 메서드
func (a *Account) Withdraw(amount float64) {
	if amount <= 0 {
		a.printer.PrintError("Invalid withdrawal amount. Must be greater than zero.")
		return
	}
	if amount > a.balance {
		a.printer.PrintError("Insufficient funds.")
		return
	}
	a.balance -= amount
	a.printer.PrintWithdraw(amount, a.balance)
}

// Balance 메서드
func (a *Account) Balance() float64 {
	a.printer.PrintBalance(a.balance)
	return a.balance
}

// Error
func (a *Account) Error(errorMessage string) {
	a.printer.PrintError(errorMessage)
	return
}

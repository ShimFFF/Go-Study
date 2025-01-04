package printer

// AccountOutput 인터페이스 정의
type AccountOutput interface {
	PrintDeposit(amount, balance float64)
	PrintWithdraw(amount, balance float64)
	PrintError(message string)
	PrintBalance(balance float64)
}

package printer

import (
	"fmt"
	"os"
)

// FilePrinter: 파일 출력 구현체
type FilePrinter struct {
	file *os.File
}

// NewFilePrinter: FilePrinter 생성자
func NewFilePrinter(filename string) (*FilePrinter, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FilePrinter{file: file}, nil
}

// PrintDeposit: 입금 기록
func (f *FilePrinter) PrintDeposit(amount, balance float64) {
	fmt.Fprintf(f.file, "Deposited: %.2f | New Balance: %.2f\n", amount, balance)
}

// PrintWithdraw: 출금 기록
func (f *FilePrinter) PrintWithdraw(amount, balance float64) {
	fmt.Fprintf(f.file, "Withdrew: %.2f | New Balance: %.2f\n", amount, balance)
}

// PrintError: 오류 기록
func (f *FilePrinter) PrintError(message string) {
	fmt.Fprintln(f.file, "Error:", message)
}

// PrintBalance: 잔액 기록
func (f *FilePrinter) PrintBalance(balance float64) {
	fmt.Fprintf(f.file, "Current Balance: %.2f\n", balance)
}

// Close: 파일 닫기
func (f *FilePrinter) Close() {
	f.file.Close()
}

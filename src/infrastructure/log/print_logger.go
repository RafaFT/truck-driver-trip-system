package log

import (
	"fmt"

	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type printLogger struct{}

func NewPrintLogger() usecase.Logger {
	return printLogger{}
}

func (l printLogger) Debug(msg string) {
	fmt.Println("Debug:", msg)
}

func (l printLogger) Info(msg string) {
	fmt.Println("Info:", msg)
}

func (l printLogger) Warning(msg string) {
	fmt.Println("Warning:", msg)
}

func (l printLogger) Error(msg string) {
	fmt.Println("Error:", msg)
}

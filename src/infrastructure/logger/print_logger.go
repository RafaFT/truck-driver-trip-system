package logger

import (
	"fmt"

	"github.com/rafaft/truck-driver-trip-system/usecase"
)

type printLogger struct{}

func NewPrintLogger() usecase.Logger {
	return printLogger{}
}

func (l printLogger) Debug(msg string) {
	fmt.Println(msg)
}

func (l printLogger) Info(msg string) {
	fmt.Println(msg)
}

func (l printLogger) Warning(msg string) {
	fmt.Println(msg)
}

func (l printLogger) Error(msg string) {
	fmt.Println(msg)
}

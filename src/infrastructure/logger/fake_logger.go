package logger

import "github.com/rafaft/truck-driver-trip-system/usecase"

type fakeLogger struct{}

func NewFakeLogger() usecase.Logger {
	return fakeLogger{}
}

func (l fakeLogger) Debug(msg string) {
}

func (l fakeLogger) Info(msg string) {
}

func (l fakeLogger) Warning(msg string) {
}

func (l fakeLogger) Error(msg string) {
}

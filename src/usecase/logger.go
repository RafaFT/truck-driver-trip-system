package usecase

type Logger interface {
	Debug(string)
	Info(string)
	Warning(string)
	Error(string)
}

type fakeLogger struct{}

func (l fakeLogger) Debug(msg string) {
}

func (l fakeLogger) Info(msg string) {
}

func (l fakeLogger) Warning(msg string) {
}

func (l fakeLogger) Error(msg string) {
}

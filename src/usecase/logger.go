package usecase

type Logger interface {
	Debug(string)
	Info(string)
	Warning(string)
	Error(string)
}

type FakeLogger struct{}

func (l FakeLogger) Debug(msg string) {
}

func (l FakeLogger) Info(msg string) {
}

func (l FakeLogger) Warning(msg string) {
}

func (l FakeLogger) Error(msg string) {
}

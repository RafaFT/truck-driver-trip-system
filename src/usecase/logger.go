package usecase

type Logger interface {
	Debug(string)
	Info(string)
	Warning(string)
	Error(string)
}

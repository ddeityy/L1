package main

/*
Реализовать паттерн «адаптер» на любом примере.
*/

import "fmt"

// У старого логгера функция Log которая
// используется в большом количестве кода
type OldLogger struct{}

func (l *OldLogger) Log(s string) {
	fmt.Println("Old logger:", s)
}

// NewLogger импортируем из библиотеки
// и не можем поменять его методы
type NewLogger struct{}

func (l *NewLogger) WriteLog(s string) {
	fmt.Println("New logger:", s)
}

func (logger *NewLogger) Log(s string) {
	logger.WriteLog(s)
}

// Чтобы избежать рефакторинга 1000+ линий кода
// создаём адаптер через интерфейс
type LoggerAdapter struct {
	*NewLogger
}

func NewLoggerAdapter(logger *NewLogger) Logger {
	return &LoggerAdapter{logger}
}

type Logger interface {
	Log(s string)
}

func main() {
	// Теперь мы можем без рефакторинга
	// продолжить использовать logger.Log()
	logger := NewLoggerAdapter(&NewLogger{})
	logger.Log("Hello")
	logger = &OldLogger{}
	logger.Log("Hello")

}

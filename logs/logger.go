package logger

import (
	"log"
	"os"
)

type LogInterface interface {
	Error(message string)
	Info(message string)
	Fatal(message string)
	Close()
}

type Loggers struct {
	InfoL  *InfoLogger
	ErrorL *ErrorLogger
	FatalL *FatalLogger
}

type Logger struct {
	logger    *log.Logger
	loggerCmd *log.Logger
	file      *os.File
}

func (l *Logger) Close() error {
	err := l.file.Close()
	return err
}

type ErrorLogger struct {
	*Logger
}

func newErrorLogger(logger *Logger) *ErrorLogger {
	logger.logger = log.New(logger.file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.loggerCmd = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &ErrorLogger{
		Logger: logger,
	}
}

func (l *ErrorLogger) Log(message string) {
	l.loggerCmd.Println(message)
	l.logger.Println(message)
}

type FatalLogger struct {
	*Logger
}

func newFatalLogger(logger *Logger) *FatalLogger {
	logger.logger = log.New(logger.file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.loggerCmd = log.New(os.Stdout, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &FatalLogger{
		Logger: logger,
	}
}

func (l *FatalLogger) Log(message string) {
	l.loggerCmd.Println(message)
	l.logger.Fatal(message)
}

type InfoLogger struct {
	*Logger
}

func newInfoLogger(logger *Logger) *InfoLogger {
	logger.logger = log.New(logger.file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.loggerCmd = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &InfoLogger{
		Logger: logger,
	}
}

func (l *InfoLogger) Log(message string) {
	l.loggerCmd.Println(message)
	l.logger.Println(message)
}

func newLogger(path string) (*Logger, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logger := &Logger{
		file: file,
	}
	return logger, err
}

func InitLoggers(infoPath, errorPath, fatalPath string) (*Loggers, error) {
	newInfo, err := newLogger(infoPath)
	if err != nil {
		panic(err.Error())
	}

	newError, err := newLogger(errorPath)
	if err != nil {
		panic(err.Error())
	}

	newFatal, err := newLogger(fatalPath)
	if err != nil {
		return nil, err
	}

	return &Loggers{
		InfoL:  newInfoLogger(newInfo),
		ErrorL: newErrorLogger(newError),
		FatalL: newFatalLogger(newFatal),
	}, nil
}

func (l *Loggers) Close() {
	err := l.InfoL.Close()
	if err != nil {
		panic(err.Error())
	}

	err = l.ErrorL.Close()
	if err != nil {
		panic(err.Error())
	}

	err = l.FatalL.Close()
	if err != nil {
		panic(err.Error())
	}
}

func (l *Loggers) Info(message string) {
	l.InfoL.Log(message)
}

func (l *Loggers) Error(message string) {
	l.ErrorL.Log(message)
}

func (l *Loggers) Fatal(message string) {
	l.FatalL.Log(message)
}

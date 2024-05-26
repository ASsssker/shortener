package logger

import (
	"github.com/fatih/color"
	"log"
	"os"
)

// CreateLogger создает логгер с заданными параметрами:
func CreateLogger(logMessage string, messageColor color.Attribute, flag int, file *os.File) *log.Logger {
	if logMessage == "" {
		log.Fatalln("No log message provided")
	}
	if messageColor <= 30 {
		messageColor = color.FgGreen
	}
	message := color.New(messageColor).Sprintf("%s \t", logMessage)
	if file == nil {
		file = os.Stdout
	}

	return log.New(file, message, flag)
}

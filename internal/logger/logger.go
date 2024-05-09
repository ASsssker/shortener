package logger

import (
	"github.com/fatih/color"
	"log"
	"os"
)

// CreateLogger создает логгер с заданными параметрами:
func CreateLogger(logMessage string, messageColor color.Attribute) *log.Logger {
	if logMessage == "" {
		log.Fatalln("No log message provided")
	}
	if messageColor <= 30 {
		messageColor = color.FgGreen
	}
	message := color.New(messageColor).Sprintf("%s \t", logMessage)
	return log.New(os.Stdout, message, log.Ldate|log.Ltime)
}

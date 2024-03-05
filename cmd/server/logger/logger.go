package logger

import (
	"log"
)

func Info(location string, message string) {
	log.Printf("(INFO) [%s] %s\n", location, message)
}

func Error(location string, message string) {
	log.Printf("(ERROR) [%s] %s\n", location, message)
}

func Warning(location string, message string) {
	log.Printf("(WARNING) [%s] %s\n", location, message)
}

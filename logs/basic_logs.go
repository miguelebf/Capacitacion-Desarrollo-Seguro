package main

import (
	"log"
)

const (
	LevelInfo  = "INFO"
	LevelError = "ERROR"
	LevelDebug = "DEBUG"
)

func logMessage(level, message string) {
	log.Printf("[%s] %s\n", level, message)
}

func main() {
	logMessage(LevelInfo, "Este es un mensaje de información")
	logMessage(LevelError, "Este es un mensaje de error")
	logMessage(LevelDebug, "Este es un mensaje de depuración")
}
package main

import (
	"encoding/json"
	"log"
	"time"
)

// LogLevel define los niveles de log
type LogLevel string

const (
	LevelInfo  LogLevel = "INFO"
	LevelError LogLevel = "ERROR"
	LevelDebug LogLevel = "DEBUG"
)

// LogEntry define la estructura de un log
type LogEntry struct {
	Timestamp string      `json:"timestamp"`
	Level     LogLevel    `json:"level"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"` // Campo opcional para datos adicionales
}

// Log es una función personalizada para generar logs
func Log(level LogLevel, message string, data interface{}) {
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339), // Formato ISO 8601
		Level:     level,
		Message:   message,
		Data:      data,
	}

	// Convertir la entrada a JSON
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Println("Error al convertir el log a JSON:", err)
		return
	}

	// Escribir el log en la salida estándar (o un archivo)
	log.Println(string(jsonEntry))
}

func main() {
	// Ejemplo de uso
	Log(LevelInfo, "La aplicación ha iniciado", nil)
	Log(LevelDebug, "Depuración de usuario", map[string]interface{}{
		"user_id": 123,
		"email":   "john.doe@example.com",
	})
	Log(LevelError, "Error al conectar a la base de datos", map[string]interface{}{
		"error":   "connection timeout",
		"timeout": 30,
	})
}
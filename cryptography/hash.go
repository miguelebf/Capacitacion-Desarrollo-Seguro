package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Función para hashear con SHA-256
func hashSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func main() {
	text := "Hola, mundo! 1234"
	hashedText := hashSHA256(text)

	fmt.Println("Texto original:", text)
	fmt.Println("Hash SHA-256:", hashedText)
}

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Función para cifrar texto con AES-GCM
func encryptAES(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12) // GCM usa un nonce de 12 bytes
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Función para descifrar texto con AES-GCM
func decryptAES(ciphertext string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, encryptedData := data[:nonceSize], data[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func main() {
	key := []byte("12345678901234567890123456789012") // Clave de 32 bytes (AES-256)
	plaintext := "Hola, mundo! 1234"

	encrypted, err := encryptAES(plaintext, key)
	if err != nil {
		fmt.Println("Error cifrando:", err)
		return
	}

	fmt.Println("Texto cifrado:", encrypted)

	decrypted, err := decryptAES(encrypted, key)
	if err != nil {
		fmt.Println("Error descifrando:", err)
		return
	}

	fmt.Println("Texto descifrado:", decrypted)
}

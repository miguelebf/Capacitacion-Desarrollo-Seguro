package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Generar una clave privada RSA (inseguro para producción)
	privateKey, err := ssh.NewSignerFromKey(generatePrivateKey())
	if err != nil {
		log.Fatalf("Error creando clave privada: %v", err)
	}

	// Configuración del servidor SSH
	config := &ssh.ServerConfig{
		NoClientAuth: true, // Inseguro: permite acceso sin autenticación
	}
	config.AddHostKey(privateKey)

	listener, err := net.Listen("tcp", "127.0.0.1:2222")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
	defer listener.Close()

	fmt.Println("Servidor SSH en ejecución en 127.0.0.1:2222")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error aceptando conexión: %v", err)
			continue
		}
		go handleConnection(conn, config)
	}
}

func generatePrivateKey() *ssh.Certificate {
	// Aquí iría la lógica de generación de una clave privada
	return nil
}

func handleConnection(conn net.Conn, config *ssh.ServerConfig) {
	ssh.NewServerConn(conn, config) // Posible vulnerabilidad
}

package main

import (
	"errors"
	"log"
	"os"
)

// User representa un usuario con información sensible
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

// getUserByID simula la obtención de un usuario por ID
func getUserByID(id int) (*User, error) {
	// Simulación de un error en la base de datos
	if id == 0 {
		return nil, errors.New("error: usuario no encontrado")
	}

	// Simulación de un usuario encontrado
	return &User{
		ID:       1,
		Username: "johndoe",
		Password: "s3cr3t",
		Email:    "john.doe@example.com",
	}, nil
}

func main() {
	userID := 0 // ID de usuario que no existe

	user, err := getUserByID(userID)
	if err != nil {
		// Registro seguro del error sin exponer información sensible
		log.Printf("Error al obtener el usuario: %v", err)
		os.Exit(1)
	}

	// Registro seguro del usuario sin exponer información sensible
	log.Printf("Usuario obtenido: ID=%d, Username=%s", user.ID, user.Username)
}
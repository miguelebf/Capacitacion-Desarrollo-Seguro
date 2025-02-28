package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// Definir la estructura con etiquetas de validación
type User struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"gte=18,lte=100"`
}

// Inicializar validador
var validate = validator.New()

func createUser(c *gin.Context) {
	var user User

	// Bind JSON al struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar manualmente con go-playground/validator
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario creado!", "user": user})
}

func main() {
	r := gin.Default()
	r.POST("/users", createUser)
	r.Run(":8080")
}
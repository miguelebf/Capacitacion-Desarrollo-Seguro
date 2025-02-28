package main

import (
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
)

func main() {
    // Cargar las variables de entorno desde el archivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error cargando el archivo .env: %v", err)
    }

    // Acceder a las variables de entorno
    dbURL := os.Getenv("DATABASE_URL")
    port := os.Getenv("PORT")
    debug := os.Getenv("DEBUG")

    // Imprimir las variables de entorno
    fmt.Println("DATABASE_URL:", dbURL)
    fmt.Println("PORT:", port)
    fmt.Println("DEBUG:", debug)
}
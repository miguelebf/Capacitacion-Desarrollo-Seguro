package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"
)

func main() {
    // Cadena de conexión a PostgreSQL
    connStr := "user=postgres dbname=testdb sslmode=disable password=postgres host=localhost"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Crear la tabla `users` si no existe
    createTableQuery := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name TEXT
        );
    `
    _, err = db.Exec(createTableQuery)
    if err != nil {
        log.Fatal(err)
    }

    // Insertar 10 registros de prueba
    names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack"}
    for _, name := range names {
        insertQuery := "INSERT INTO users (name) VALUES ($1)"
        _, err := db.Exec(insertQuery, name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Inserted: %s\n", name)
    }

    fmt.Println("Database populated with 10 records!")
}
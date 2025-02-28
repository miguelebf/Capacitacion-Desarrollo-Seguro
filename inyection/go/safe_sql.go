package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    _ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
    var err error
    connStr := "user=postgres dbname=testdb sslmode=disable password=postgres host=localhost"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
}

func safeHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("id")
    query := "SELECT name FROM users WHERE id = $1"

    var name string
    err := db.QueryRow(query, userID).Scan(&name)
    if err != nil {
        http.Error(w, "Error querying database", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "User: %s", name)
}

func main() {
    initDB()
    http.HandleFunc("/safe", safeHandler)
    log.Println("Server Inyection Good started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
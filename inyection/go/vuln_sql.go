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

func vulnerableHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("id")
    query := fmt.Sprintf("SELECT name FROM users WHERE id = %s", userID)
    //"SELECT name FROM users WHERE id = 1 OR 1=1"
    rows, err := db.Query(query)
    if err != nil {
        http.Error(w, "Error querying database", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var names []string
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            http.Error(w, "Error scanning row", http.StatusInternalServerError)
            return
        }
        names = append(names, name)
    }

    fmt.Fprintf(w, "Names: %v", names)
}

func main() {
    initDB()
    http.HandleFunc("/vulnerable", vulnerableHandler)
    log.Println("Server started at SQL Inyection :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
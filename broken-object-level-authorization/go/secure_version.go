package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = map[int]User{
	1: {ID: 1, Username: "user1", Email: "user1@example.com"},
	2: {ID: 2, Username: "user2", Email: "user2@example.com"},
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// Read Header x-user-id simulating user authentication
	authenticatedUserIDStr := r.Header.Get("x-user-id")
	authenticatedUserID, err := strconv.Atoi(authenticatedUserIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userIDStr := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// ------------ Verificación de autorización ------------
	if userID != authenticatedUserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, exists := users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
        http.Error(w, "Failed to encode user data", http.StatusInternalServerError)
    }
}

func main() {
	http.HandleFunc("/user", getUser)
	fmt.Println("Go Secure Version Server Broken Object Level Authorization started at :8080")
	server := &http.Server{
        Addr:         ":8080",
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  15 * time.Second,
    }
	if err := server.ListenAndServe(); err != nil {
        fmt.Printf("Server failed to start")
    }
}
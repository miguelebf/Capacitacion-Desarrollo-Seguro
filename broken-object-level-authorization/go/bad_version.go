package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

	// ------------ Mala Verificación de autorización ------------
	if authenticatedUserID != authenticatedUserID {
	// To-do

	}

	userIDStr := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, exists := users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/user", getUser)
	fmt.Println("Bad Version Server Broken Object Level Authorization started at :8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lestrrat-go/jwx/jwe"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/lestrrat-go/jwx/jwa"
)

var rsaPrivateKey, _ = rsa.GenerateKey(rand.Reader, 2048)
var rsaPublicKey = &rsaPrivateKey.PublicKey

type Claims struct {
	Username string `json:"username"`
	ExpiresAt int64 `json:"exp"`
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	claims := &Claims{
		Username: username,
		ExpiresAt: expirationTime,
	}

	token := jwt.New()
	token.Set("username", claims.Username)
	token.Set("exp", claims.ExpiresAt)

	payload, err := jwt.Sign(token, jwa.HS256, []byte("your_secret_key"))
	if err != nil {
		return "", err
	}

	encrypted, err := jwe.Encrypt(payload, jwa.RSA_OAEP, rsaPublicKey, jwa.A128GCM, jwa.NoCompress)
	if err != nil {
		return "", err
	}

	return string(encrypted), nil
}

func ValidateToken(tknStr string) (*Claims, error) {
	decrypted, err := jwe.Decrypt([]byte(tknStr), jwa.RSA_OAEP, rsaPrivateKey)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(decrypted, jwt.WithVerify(jwa.HS256, []byte("your_secret_key")))
	if err != nil {
		return nil, err
	}

	claims := &Claims{
		Username: token.Subject(),
		ExpiresAt: token.Expiration().Unix(),
	}

	return claims, nil
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	// For demonstration, we're assuming "admin"/"password" as valid credentials.
	// Simulating user authentication
	username_database := "admin"
	password_database := "password"
	fmt.Println("Username:", username)
    fmt.Println("Password:", password)
	if username != username_database || password != password_database {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := GenerateToken(username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour),
	})
}

func handleProtectedRoute(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No token", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tokenStr := c.Value
	claims, err := ValidateToken(tokenStr)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s!", claims.Username)))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", handleLogin).Methods("POST")
	r.HandleFunc("/protected", handleProtectedRoute).Methods("GET")
	fmt.Println("JWE Server One Route started at :8080")
	http.ListenAndServe(":8080", r)
}

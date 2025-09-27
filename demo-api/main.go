package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var keycloakIssuer = "http://localhost:8080/realms/demo"
var keycloakCertsURL = keycloakIssuer + "/protocol/openid-connect/certs"

// Middleware to validate JWT
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse without verification (for demo), normally use JWKS from certsURL
		token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Fprintf(w, "Hello %s! Token issued by %s\n",
				claims["preferred_username"], claims["iss"])
		} else {
			http.Error(w, "invalid claims", http.StatusUnauthorized)
		}
	})
}

func main() {
	http.Handle("/secure", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You accessed a secure endpoint!"))
	})))

	log.Println("Server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

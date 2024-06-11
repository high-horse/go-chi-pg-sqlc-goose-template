package middleware

import (
	"context"
	"net/http"
	"server/http/helper"
)

// JWTMiddleware is a middleware that checks for a valid JWT cookie,
// extracts user info from the token, and stores it in the context.
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the "jwt" cookie
		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "Missing or invalid JWT cookie", http.StatusUnauthorized)
			return
		}

		// Verify the token and extract claims
		tokenString := cookie.Value
		claims, err := helper.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired JWT token", http.StatusUnauthorized)
			return
		}

		// Extract user info from the claims
		email, ok := claims["email"].(string)
		if !ok {
			http.Error(w, "Invalid token payload1", http.StatusUnauthorized)
			return
		}

		userInfo := helper.UserInfo{
			Email:    email,
			Username: claims["username"].(string),
			ID:       int64(claims["id"].(float64)),
			// Add more fields as needed
		}

		// Store the user info in the context
		ctx := context.WithValue(r.Context(), "user", userInfo)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

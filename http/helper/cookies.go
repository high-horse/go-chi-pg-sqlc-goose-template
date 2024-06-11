package helper

import (
	"net/http"
	"time"
)


// Function to set the cookie in a Chi handler
func SetJWTToken(w http.ResponseWriter, name string, token string) {
    // Create the cookie with equivalent properties
    cookie := &http.Cookie{
        Name:     "jwt",
        Value:    token,
        // Expires is calculated as current time + 3600 seconds (1 hour)
        Expires:  time.Now().Add(3600 * time.Second),
        Path:     "/",
        Domain:   "localhost",  // Optional, set if needed
        HttpOnly: true,
        Secure:   false,        // false means the cookie is sent over HTTP as well
    }

    // Set the cookie in the response
    http.SetCookie(w, cookie)
}

func UnsetJWTToken(w http.ResponseWriter, name string) {
	cookie := &http.Cookie{
		Name:     name,           // The name of the cookie to unset
		Value:    "",             // Empty value
		Expires:  time.Now().Add(-1 * time.Hour), // Expired date to remove the cookie
		Path:     "/",            // The path should match the original cookie path
		Domain:   "localhost",    // Domain restriction should match the original cookie domain
		HttpOnly: true,           // HttpOnly to prevent JavaScript access
		Secure:   false,          // Not restricted to HTTPS (set to true in production if using HTTPS)
	}
	http.SetCookie(w, cookie)    // Attach the expired cookie to the response to remove it
}
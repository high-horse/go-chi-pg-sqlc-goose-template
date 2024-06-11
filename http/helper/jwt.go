package helper

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))


func CreateToken(id int64,email string, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"id":    id,
			"email": email,
			"username": username,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}


func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid or expired JWT token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid JWT claims")
	}

	return claims, nil
}


type UserInfo struct {
	ID int64
	Email string
	Username string
}

// GetUserFromContext retrieves the UserInfo stored in the context.
func GetUserFromContext(ctx context.Context) (*UserInfo, bool) {
	userInfo, ok := ctx.Value("user").(UserInfo)
	return &userInfo, ok
}
package middleware

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)
var secretKey = []byte("khadde")
func JWT_Token(username string, password string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
			"username": username,
			"password": password,
			"exp": time.Now().Add(time.Hour * 720).Unix(),
        })
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
		return "", err
    }
	
	return tokenString, err
}
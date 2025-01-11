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

// check, err := JWT_Authenticate("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkxODYyMjEsInVzZXJuYW1lIjoibWF5dXIifQ.KbwHfnPtzmT7hrk3PwG6BYfXOWUs99V3JiOhPLX2jmA")
// if err != nil {
// 	log.Error(err)
// }
// fmt.Println(check)
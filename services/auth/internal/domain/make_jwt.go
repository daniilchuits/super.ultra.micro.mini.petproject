package domain

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func MakingJWT(id int, secret string) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp":     time.Now().Add(2 * time.Hour),
			"user_id": id,
		})
	jwt, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("Making jwt err:", err)
	}
	return jwt, err
}

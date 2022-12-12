package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(userId int, first_name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["first_name"] = first_name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

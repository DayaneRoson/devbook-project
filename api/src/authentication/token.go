package authentication

import (
	"api/src/config"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

// CreateToken returns a signed token with user permissions
func CreateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

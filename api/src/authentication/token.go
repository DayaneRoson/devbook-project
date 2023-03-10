package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

// CreateToken returns a signed token with user permissions
func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken validates a token received
func ValidateToken(r *http.Request) error {
	stringToken := extractToken(r)
	token, error := jwt.Parse(stringToken, returnKey)
	if error != nil {
		return error
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("Invalid Token")
}

func ExtractUserId(r *http.Request) (uint64, error) {
	stringToken := extractToken(r)
	token, error := jwt.Parse(stringToken, returnKey)
	if error != nil {
		return 0, error
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, error := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if error != nil {
			return 0, error
		}
		return userId, nil
	}
	return 0, errors.New("Invalid Token")

}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected Signing Method %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

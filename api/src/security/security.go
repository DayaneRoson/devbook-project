package security

import "golang.org/x/crypto/bcrypt"

//Hash receives a string value and hashes it
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerifyPassword compares a password and a hash and returns if they match
func VerifyPassword(hashPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(inputPassword))
}

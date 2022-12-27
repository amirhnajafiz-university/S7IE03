package crypto

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	errEmptyPassword = errors.New("password cannot be empty")
)

// HashData gets a data and returns a hash string.
func HashData(data string) (string, error) {
	if len(data) == 0 {
		return "", errEmptyPassword
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)

	return string(hash), err
}

// IsEqual checks the equality of hash value to its origin.
func IsEqual(data string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(data), []byte(hash)) == nil
}

package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	Key    string
	Expire int
}

// GenerateJWT creates a new JWT token.
func (a *Auth) GenerateJWT(username string, password string) (string, error) {
	// create a new token
	token := jwt.New(jwt.SigningMethodEdDSA)

	// create claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Duration(a.Expire) * time.Minute)
	claims["username"] = username
	claims["password"] = password

	// generate token string
	tokenString, err := token.SignedString(a.Key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

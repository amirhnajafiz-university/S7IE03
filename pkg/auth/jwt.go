package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	errSigningMethod = errors.New("error in signing method")
	errInvalidToken  = errors.New("token is invalid")
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

// ParseJWT gets a token string and extracts the data.
func (a *Auth) ParseJWT(tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return "", errSigningMethod
		}

		return a.Key, nil
	})
	if err != nil {
		return "", "", err
	}

	// taking out claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		password := claims["password"].(string)

		return username, password, nil
	}

	return "", "", errInvalidToken
}

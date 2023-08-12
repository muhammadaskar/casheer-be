package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type JWTAuthentication interface {
	GenerateToken(userID int, email string, role int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtAuth struct {
}

var SECRET_KEY = []byte("s3cr3T_k3Y")

func NewJWTAuth() *jwtAuth {
	return &jwtAuth{}
}

func (j *jwtAuth) GenerateToken(userID int, email string, role int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["email"] = email
	claim["role"] = role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (j *jwtAuth) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

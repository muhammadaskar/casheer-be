package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTAuthentication interface {
	GenerateToken(userID int, email string, role int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtAuth struct {
}

type customToken struct {
	UserID int    `json:"id"`
	Email  string `json:"email"`
	Role   int    `json:"role"`
	jwt.StandardClaims
}

var SECRET_KEY = []byte("s3cr3T_k3Y")

func NewJWTAuth() *jwtAuth {
	return &jwtAuth{}
}

func (j *jwtAuth) GenerateToken(userID int, email string, role int) (string, error) {
	exp := time.Now().Add(24 * time.Hour).Unix()
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = exp
	// claims := &customToken{
	// 	UserID: userID,
	// 	Email:  email,
	// 	Role:   role,
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: exp,
	// 	},
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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

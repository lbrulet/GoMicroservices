package handler

import (
	"github.com/dgrijalva/jwt-go"
)

const SIGNED_KEY = "pingouin123"

type JwtService struct {
	SignedKey string
}

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

func NewJwtService(signedKey string) *JwtService {
	return &JwtService{
		SignedKey: signedKey,
	}
}

func (t *JwtService) CreateToken(id int) (string, error) {
	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(t.SignedKey))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (t *JwtService) VerifyToken(tokenToValidate string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenToValidate, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.SignedKey), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Id, nil
	}
	return -1, err
}

package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtService struct {
	secretKey string
	issuer string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: "secret-key",
		issuer:    "book-api",
	}
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GeraToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour*2).Unix(),
			Issuer: s.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidaToken(token string) bool {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid{
			return nil, fmt.Errorf("token inválido: %v", token)
		}
		return []byte(s.secretKey), nil
	})

	return err == nil
}
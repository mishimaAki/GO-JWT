package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"GO-JWT/internal/domain/model"
)

type jwtAuth struct {
	secretKey []byte
}

func NewJWTAuth(secretKey string) *jwtAuth {
	return &jwtAuth{
		secretKey: []byte(secretKey),
	}
}

func (j *jwtAuth) GenerateToken(userID uint, role string) (string, error) {
	claims := &model.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: userID,
		Role:   role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

package model

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	jwt.RegisteredClaims
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
}

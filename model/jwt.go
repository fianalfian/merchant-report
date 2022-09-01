package model

import "github.com/golang-jwt/jwt/v4"

type JWTPayload struct {
	UserID     uint64 `json:"user_id"`
	MerchantID uint64 `json:"merchant_id"`
	jwt.RegisteredClaims
}

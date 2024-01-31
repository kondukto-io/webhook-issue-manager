package handler

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/webhook-issue-manager/model"
)

var (
	secretKey       = "secretKey"
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired Token")
)

func verifyToken(token string) (*model.Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &model.Payload{}, keyFunc)
	if err != nil {
		return nil, ErrInvalidToken
	} else if claims, ok := jwtToken.Claims.(*model.Payload); ok {
		return claims, nil
	} else {
		return nil, ErrInvalidToken
	}
}

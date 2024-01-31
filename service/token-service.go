package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/webhook-issue-manager/model"
	tokenrepository "github.com/webhook-issue-manager/storage/token-repository"
)

var (
	secretKey = "secretKey"
)

type TokenService interface {
	CreateToken() (*model.Token, error)
	GetToken(tokenId string) (*model.Token, error)
}

type tokenService struct{}

var (
	tokenRepo = tokenrepository.NewTokenRepository()
)

func NewTokenService() TokenService {
	return &tokenService{}
}

func (*tokenService) CreateToken() (*model.Token, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	claims := jwt.MapClaims{
		"id":         tokenID,
		"issued_at":  time.Now(),
		"expired_at": time.Now().Add(time.Hour * 24 * 7),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	token := model.Token{TokenID: tokenID.String(), TokenStr: tokenStr}

	addedToken, err := tokenRepo.AddToken(&token)
	if err != nil {
		return nil, err
	}

	return addedToken, nil
}

func (*tokenService) GetToken(tokenId string) (*model.Token, error) {
	token, err := tokenRepo.GetToken(tokenId)
	if err != nil {
		return nil, err
	}
	return token, nil
}

package tokenrepository

import (
	"errors"
	"fmt"

	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/storage/postgres"
)

type TokenRepository interface {
	AddToken(token *model.Token) (*model.Token, error)
	GetToken(tokenId string) (*model.Token, error)
}

type tokenRepository struct{}

func NewTokenRepository() TokenRepository {
	return &tokenRepository{}
}

// AddToken implements TokenRepository
func (*tokenRepository) AddToken(token *model.Token) (*model.Token, error) {
	db := postgres.Init()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	err := db.Create(token).Error
	if err != nil {
		return nil, err
	}
	return token, nil
}

// GetToken implements TokenRepository
func (*tokenRepository) GetToken(tokenId string) (*model.Token, error) {
	var token model.Token
	db := postgres.Init()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()
	if tokenId == "" {
		fmt.Println("TokenID can not be empty")
	}
	result := db.Where("token_id = ?", tokenId).Find(&token)
	if result.Error != nil {
		return nil, errors.New("record is not found")
	}
	return &token, nil
}

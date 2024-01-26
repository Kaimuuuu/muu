package token

import (
	"kaimuu/model"

	"github.com/google/uuid"
)

func NewTokenService(tokenRepo TokenRepository, promotionRepo PromotionRepository) *TokenService {
	return &TokenService{
		tokenRepo:     tokenRepo,
		promotionRepo: promotionRepo,
	}
}

type TokenService struct {
	tokenRepo     TokenRepository
	promotionRepo PromotionRepository
}

type TokenRepository interface {
	Get(token string) (*model.Client, error)
	All() ([]model.Client, error)
	Delete(token string) error
	Insert(c *model.Client) error
}

type PromotionRepository interface {
	GetById(id string) (*model.Promotion, error)
}

func GenerateToken() string {
	return uuid.NewString()
}

type GenerateTokenRequest struct {
	PromotionId string `validate:"required" json:"promotionId"`
	TableNumber int8   `validate:"required,number" json:"tableNumber"`
	Size        int8   `validate:"min=1" json:"size"`
}

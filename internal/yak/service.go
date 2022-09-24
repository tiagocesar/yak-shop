package yak

import (
	"context"

	"github.com/tiagocesar/yak-shop/internal/models"
)

type Service struct {
	herd []models.Yak // In a real scenario it would be a data store
}

func NewService(herd []models.Yak) *Service {
	return &Service{herd: herd}
}

func (s *Service) Process(ctx context.Context, day int) (int, int, error) {
	panic("implement me")
}

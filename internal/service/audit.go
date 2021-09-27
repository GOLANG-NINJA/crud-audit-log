package service

import (
	"context"

	"github.com/GOLANG-NINJA/crud-audit-log/internal/domain"
)

type Repository interface {
	Insert(ctx context.Context, item domain.LogItem) error
}

type Audit struct {
	repo Repository
}

func NewAudit(repo Repository) *Audit {
	return &Audit{
		repo: repo,
	}
}

func (s *Audit) Insert(ctx context.Context, req *domain.LogRequest) error {
	item := domain.LogItem{
		Action:    req.GetEntity().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return s.repo.Insert(ctx, item)
}

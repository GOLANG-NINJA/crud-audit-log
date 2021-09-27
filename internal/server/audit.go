package server

import (
	"context"

	"github.com/GOLANG-NINJA/crud-audit-log/internal/domain"
)

type AuditService interface {
	Insert(ctx context.Context, req *domain.LogRequest) error
}

type AuditServer struct {
	service AuditService
}

func NewAuditServer(service AuditService) *AuditServer {
	return &AuditServer{
		service: service,
	}
}

func (h *AuditServer) Log(ctx context.Context, req *domain.LogRequest) (*domain.Empty, error) {
	err := h.service.Insert(ctx, req)

	return nil, err
}

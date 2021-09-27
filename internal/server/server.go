package server

import (
	"fmt"
	"net"

	"github.com/GOLANG-NINJA/crud-audit-log/internal/domain"
	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv     *grpc.Server
	auditServer domain.AuditServiceServer
}

func New(auditServer domain.AuditServiceServer) *Server {
	return &Server{
		grpcSrv:     grpc.NewServer(),
		auditServer: auditServer,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	domain.RegisterAuditServiceServer(s.grpcSrv, s.auditServer)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}

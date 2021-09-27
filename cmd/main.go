package main

import (
	"fmt"
	"log"
	"time"

	"github.com/GOLANG-NINJA/crud-audit-log/internal/repository"
	"github.com/GOLANG-NINJA/crud-audit-log/internal/server"
	"github.com/GOLANG-NINJA/crud-audit-log/internal/service"
)

func main() {
	auditRepo := &repository.Audit{}
	auditService := service.NewAudit(auditRepo)
	auditSrv := server.NewAuditServer(auditService)
	srv := server.New(auditSrv)

	fmt.Println("SERVER STARTED", time.Now())

	if err := srv.ListenAndServe(9000); err != nil {
		log.Fatal(err)
	}
}

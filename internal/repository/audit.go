package repository

import (
	"context"
	"fmt"

	audit "github.com/GOLANG-NINJA/crud-audit-log/pkg/domain"
)

type Audit struct {
}

func (r *Audit) Insert(ctx context.Context, item audit.LogItem) error {
	fmt.Printf("INSERTED %+v\n", item)
	return nil
}

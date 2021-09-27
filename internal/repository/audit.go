package repository

import (
	"context"
	"fmt"

	"github.com/GOLANG-NINJA/crud-audit-log/internal/domain"
)

type Audit struct {
}

func (r *Audit) Insert(ctx context.Context, item domain.LogItem) error {
	fmt.Printf("INSERTED %+v\n", item)
	return nil
}

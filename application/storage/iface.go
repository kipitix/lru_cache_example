package storage

import (
	"context"
	"time"
)

type Storage interface {
	Get(ctx context.Context, key string, template interface{}) error
	Set(ctx context.Context, key string, value interface{}, dataTTL time.Duration) error
}

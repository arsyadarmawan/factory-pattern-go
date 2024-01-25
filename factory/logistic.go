package factory

import (
	"context"
	"time"
)

type LogisticFactory interface {
	Select(logisticDelivery string) LogisticProduct
}

type LogisticProduct interface {
	SetDuration(ctx context.Context) time.Duration
	Trip(ctx context.Context) string
}

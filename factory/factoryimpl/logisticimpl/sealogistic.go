package logisticimpl

import (
	"context"
	"time"
)

type SeaLogisticFactoryOpts struct {
}

type SeaLogistic struct {
	opts SeaLogisticFactoryOpts
}

func NewSeaLogistic(options SeaLogisticFactoryOpts) *SeaLogistic {
	return &SeaLogistic{
		opts: options,
	}
}

func (s SeaLogistic) SetDuration(ctx context.Context) time.Duration {
	duration := time.Hour * 5
	return time.Duration(duration)
}

func (s SeaLogistic) Trip(ctx context.Context) string {
	return "Boat"
}

package logisticimpl

import (
	"context"
	"time"
)

type RoadLogisticFactoryOpts struct {
}

type RoadLogistic struct {
	opts RoadLogisticFactoryOpts
}

func NewRoadLogistic(opts RoadLogisticFactoryOpts) *RoadLogistic {
	return &RoadLogistic{
		opts: opts,
	}
}

func (r RoadLogistic) SetDuration(ctx context.Context) time.Duration {
	duration := time.Hour * 3
	return time.Duration(duration)
}

func (r RoadLogistic) Trip(ctx context.Context) string {
	return "land"
}

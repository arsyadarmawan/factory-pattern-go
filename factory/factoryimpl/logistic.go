package factoryimpl

import (
	"factory/factory"
	"factory/factory/factoryimpl/logisticimpl"
)

type LogisticFactoryOpts struct {
}

type LogisticFactory struct {
	roadLogistic *logisticimpl.RoadLogistic
	sealogistic  *logisticimpl.SeaLogistic
}

func NewLogisticFactory() *LogisticFactory {
	roadLogisticImpl := logisticimpl.NewRoadLogistic(logisticimpl.RoadLogisticFactoryOpts{})
	seaLogisticImpl := logisticimpl.NewSeaLogistic(logisticimpl.SeaLogisticFactoryOpts{})
	return &LogisticFactory{
		roadLogistic: roadLogisticImpl,
		sealogistic:  seaLogisticImpl,
	}
}

func (l LogisticFactory) Select(logisticDelivery string) factory.LogisticProduct {
	switch logisticDelivery {
	case Sea:
		return l.sealogistic
	default:
		return l.roadLogistic
	}
}

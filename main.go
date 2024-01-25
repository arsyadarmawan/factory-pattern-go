package main

import (
	"context"
	"factory/factory/factoryimpl"
	"fmt"
)

func main() {
	ctx := context.Background()
	// factory road logistic
	logistic := factoryimpl.NewLogisticFactory()
	road := logistic.Select("road")
	tripRoad := road.Trip(ctx)
	tripDuration := road.SetDuration(ctx)
	fmt.Println(tripRoad)
	fmt.Println(tripDuration)
	fmt.Println("----------------")
	// factory sea logistic

	sea := logistic.Select("sea")
	tripSea := sea.Trip(ctx)
	tripSeaDuration := sea.SetDuration(ctx)
	fmt.Println(tripSea)
	fmt.Println(tripSeaDuration)
}

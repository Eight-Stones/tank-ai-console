package entity

import (
	"github.com/Eight-Stones/ecs-tank-engine/v2/components"
)

type Direction components.Direction

func (d Direction) String() string {
	return components.Direction(d).String()
}

type Cell struct {
	X         int
	Y         int
	Direction string
	Type      string
}

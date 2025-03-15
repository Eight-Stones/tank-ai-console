package entity

import (
	"context"
)

type Bot interface {
	Run(ctx context.Context)
	Stop()
}

type Player struct {
	ID   string
	Game string
	Name string
	Tank Tank
}

type Action struct {
	ID   int
	Type string
}

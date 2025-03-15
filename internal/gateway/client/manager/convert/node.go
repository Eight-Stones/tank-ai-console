package convert

import (
	"go-micro-service-template/entity"
	"go-micro-service-template/internal/gateway/client/manager/model"
)

func ModelToPlayer(in *model.Player) *entity.Player {
	return &entity.Player{
		ID:   in.ID,
		Name: in.Name,
		Tank: entity.Tank{
			Object: entity.Object{
				ID:      in.ID,
				HP:      in.HitPoints,
				IsAlive: in.IsAlive,
				Coordinates: entity.Coordinates{
					X: in.X,
					Y: in.Y,
				},
				Direction: in.Direction,
			},
			Ammo: in.Ammo,
		},
	}
}

func ModelToPlayers(in []*model.Player) []*entity.Player {
	out := make([]*entity.Player, len(in))
	for i := range in {
		out[i] = ModelToPlayer(in[i])
	}
	return out
}

func ModelToGame(in *model.Game) *entity.Game {
	return &entity.Game{
		ID:     in.ID,
		Status: in.Status,
		Player: ModelToPlayers(in.Players),
	}
}

func ModelToGames(in []*model.Game) []*entity.Game {
	out := make([]*entity.Game, 0, len(in))
	for i := range in {
		out = append(out, ModelToGame(in[i]))
	}
	return out
}

func ModelToCell(in *model.GameCell) *entity.Cell {
	return &entity.Cell{
		X:         in.X,
		Y:         in.Y,
		Direction: in.Direction,
		Type:      in.Type,
	}
}

func ModelToColumn(in model.GameColumn) []*entity.Cell {
	out := make([]*entity.Cell, 0, len([]*model.GameCell{}))
	for i := range in {
		out = append(out, ModelToCell(&in[i]))
	}
	return out
}

func ModelToCells(in model.GameField) [][]*entity.Cell {
	out := make([][]*entity.Cell, 0, len(in))
	for i := range in {
		out = append(out, ModelToColumn(in[i]))
	}
	return out
}

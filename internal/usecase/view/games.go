package view

import (
	"context"

	"go-micro-service-template/entity"
	er "go-micro-service-template/pkg/error"
)

func (v *View) readGames(ctx context.Context) error {
	games, err := v.manager.Games(ctx)
	if err != nil {
		return er.Wrap(err, "Games")
	}

	v.updateGames(games)

	if len(games) > 0 {
		v.app.SetFocus(v.games)
	}

	return nil
}

func (v *View) justReadGames(ctx context.Context) error {
	games, err := v.manager.Games(ctx)
	if err != nil {
		return er.Wrap(err, "Games")
	}

	v.updateGames(games)

	return nil
}

func (v *View) updateGames(in []*entity.Game) {
	v.games.ShowSecondaryText(true)
	count := v.games.GetItemCount()

	if len(in) == 0 {
		return
	}

	if count > 1 {
		for i := 1; i < count; i++ {
			v.games.RemoveItem(i)
		}
	}

	for idx, g := range in {
		v.games.AddItem(g.ID, g.Status, keyboardRunes[idx], func() {
			id := g.ID
			v.selectGame(id)
		})
	}
}

package view

import (
	"context"
	"fmt"

	"github.com/rivo/tview"

	"go-micro-service-template/entity"
)

func (v *View) selectGame(id string) {
	v.game.setID(id)
	v.readGame(id)

	g, err := v.manager.Game(context.Background(), id)
	if err != nil {
		return
	}
	if g.Status == "Processing" {
		v.jobs.start <- g.ID
	}

	v.app.SetFocus(v.gameControls)
}

func (v *View) readGame(id string) {
	g, err := v.manager.Game(context.Background(), id)
	if err != nil {
		return
	}
	v.updateGame(g)
}

func (v *View) updateGame(in *entity.Game) {
	v.players.Clear()
	v.prepareGame()

	for idx, p := range in.Player {
		name := tview.NewTableCell(p.Name).SetAlign(tview.AlignCenter)
		hpval := "0"
		if p.Tank.HP > 0 {
			hpval = fmt.Sprintf("%v", p.Tank.HP)
		}
		hp := tview.NewTableCell(hpval).SetAlign(tview.AlignCenter)
		ammo := tview.NewTableCell(fmt.Sprintf("%v", p.Tank.Ammo)).SetAlign(tview.AlignCenter)
		x := tview.NewTableCell(fmt.Sprintf("%d", p.Tank.X)).SetAlign(tview.AlignCenter)
		y := tview.NewTableCell(fmt.Sprintf("%d", p.Tank.Y)).SetAlign(tview.AlignCenter)
		direction := tview.NewTableCell(p.Tank.Direction).SetAlign(tview.AlignCenter)
		id := idx + 1
		v.players.SetCell(id, 0, name)
		v.players.SetCell(id, 1, hp)
		v.players.SetCell(id, 2, ammo)
		v.players.SetCell(id, 3, x)
		v.players.SetCell(id, 4, y)
		v.players.SetCell(id, 5, direction)
	}
}

func (v *View) runGame(id string) {
	if err := v.manager.StartGame(context.Background(), id); err != nil {
		return
	}

	if err := v.justReadGames(context.Background()); err != nil {
		return
	}

	v.app.SetFocus(v.gameControls)

	//  send in job signal to reading game
	v.jobs.start <- id
}

func (v *View) stopGame(id string) {
	if err := v.manager.StopGame(context.Background(), id); err != nil {
		return
	}

	if err := v.justReadGames(context.Background()); err != nil {
		return
	}

	v.app.SetFocus(v.gameControls)

	// send in job signal to stop reading game
	v.jobs.end <- struct{}{}
}

func (v *View) addBot(id string) {
	if _, err := v.manager.AddBot(context.Background(), id); err != nil {
		return
	}
	v.readGame(id)
}

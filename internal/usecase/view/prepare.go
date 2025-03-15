package view

import (
	"context"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (v *View) prepareControls() {
	v.controls.ShowSecondaryText(false)
	v.controls.AddItem("create", "", 'q', func() {
		if err := v.createGame(context.Background()); err != nil {
			return
		}
	})
	v.controls.AddItem("get games", "", 'q', func() {
		if err := v.readGames(context.Background()); err != nil {
			return
		}
	})
	v.controls.AddItem("exit", "", 'e', func() {
		v.app.Stop()
	})
}

func (v *View) prepareGames() {
	v.games.ShowSecondaryText(true)
	v.games.AddItem("...", "", 'q', func() {
		v.app.SetFocus(v.controls)
	})
}

func (v *View) prepareGameControls() {
	v.gameControls.ShowSecondaryText(false)
	v.gameControls.AddItem("...", "", 'q', func() {
		v.app.SetFocus(v.games)
	})
	v.gameControls.AddItem("run", "", keyboardRunes[0], func() {
		v.runGame(v.game.getID())
	})
	v.gameControls.AddItem("stop", "", keyboardRunes[1], func() {
		v.stopGame(v.game.getID())
	})
	v.gameControls.AddItem("add bot", "", keyboardRunes[2], func() {
		v.addBot(v.game.getID())
	})

}

func (v *View) prepareGame() {
	v.field.SetFixed(5, 5)
	cell := tview.NewTableCell("name").SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorOrange)
	v.players.SetCell(0, 0, cell)

	cell = tview.NewTableCell(" hp ").SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorOrange)
	v.players.SetCell(0, 1, cell)

	cell = tview.NewTableCell("ammo").SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorOrange)
	v.players.SetCell(0, 2, cell)

	cell = tview.NewTableCell(" x ").SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorOrange)
	v.players.SetCell(0, 3, cell)

	cell = tview.NewTableCell(" y ").SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorOrange)
	v.players.SetCell(0, 4, cell)

	cell = tview.NewTableCell(" direction ").SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorOrange)
	v.players.SetCell(0, 5, cell)
}

func (v *View) prepareField() {
	v.field.SetFixed(15, 15)
	cell := tview.NewTableCell(" x\\y ").SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorOrange)
	v.field.SetCell(15, 15, cell)

	for jdx := range 15 {
		label := fmt.Sprintf("  %d ", jdx)
		if jdx > 9 {
			label = fmt.Sprintf(" %d ", jdx)
		}
		cell = tview.NewTableCell(label).SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorOrange)
		v.field.SetCell(reverseIDx[jdx], 15, cell)
	}

	for idx := range 15 {
		label := fmt.Sprintf("  %d ", idx)
		if idx > 9 {
			label = fmt.Sprintf(" %d ", idx)
		}
		cell = tview.NewTableCell(label).SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorOrange)
		v.field.SetCell(15, idx, cell)
	}
}

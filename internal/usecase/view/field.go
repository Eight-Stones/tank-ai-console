package view

import (
	"context"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var reverseIDx = []int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

func defineDirection(in string) string {
	switch in {
	case "up":
		return tankUpLabel
	case "down":
		return tankDownLabel
	case "left":
		return tankLeftLabel
	case "right":
		return tankRightLabel
	}
	return ""
}

func (v *View) readView(id string) {
	cells, err := v.manager.View(context.Background(), id)
	if err != nil {
		return
	}

	for idx := range cells {
		for jdx := range cells[idx] {
			label := emptyLabel
			switch cells[idx][jdx].Type {
			case "tank":
				label = defineDirection(cells[idx][jdx].Direction)
			case "bullet":
				label = bulletLabel
			}
			cell := tview.NewTableCell(label).SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite)
			v.field.SetCell(reverseIDx[cells[idx][jdx].Y], cells[idx][jdx].X, cell)
		}
	}

	return
}

package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func MakeInnerLayout() *tview.Grid {

	Grid := tview.NewGrid()

	rows := []int{-1, -2, -1}
	columns := []int{-1, -5, -1}

	Grid.SetRows(rows...)
	Grid.SetColumns(columns...)
	Grid.SetBorders(true).SetBordersColor(tcell.ColorWhite)

	for i := range len(rows) {
		for j := range len(columns) {
			if i == 1 && j == 1 {
				Grid.AddItem(
					MakeBox(),
					i,    // row
					j,    // column
					1,    // rowSpan
					1,    // colSpan
					0,    // minGridHeight
					0,    // minGridWidth
					true, // focus
				)
			}
		}
	}
	return Grid
}

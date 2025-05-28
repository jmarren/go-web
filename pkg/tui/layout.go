package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GridItem struct {
	p       tview.Primitive
	row     int
	col     int
	rowSpan int
	colSpan int
	focus   bool
}

type MyGrid struct {
	*tview.Grid
}

func (g *MyGrid) Add(gi *GridItem) {
	g.AddItem(gi.p, gi.row, gi.col, gi.rowSpan, gi.colSpan, 0, 0, gi.focus)
}

func newMyGrid() *MyGrid {
	return &MyGrid{
		tview.NewGrid(),
	}
}

func MakeLayout() (*tview.Grid, tview.Primitive) {
	Grid := &MyGrid{
		tview.NewGrid(),
	}

	rows := []int{-1, -10, -1}
	columns := []int{-1, -10, -1}

	Grid.SetRows(rows...)
	Grid.SetColumns(columns...)

	innerGrid := newMyGrid()
	innerGrid.SetRows()
	innerGrid.SetColumns(-100, -100)
	innerGrid.SetTitle(" Go Web! ")
	innerGrid.SetBorder(true).SetBorderColor(tcell.ColorWhite)

	innerLeft := newMyGrid()
	innerLeft.SetRows(-1, -1)
	innerLeft.SetColumns(-1, -50, -1)
	innerLeft.SetBackgroundColor(tcell.ColorBlack)

	innerLeft.Add(
		&GridItem{
			p:       NewBlackBox(),
			row:     0,
			col:     0,
			rowSpan: 1,
			colSpan: 2,
		})

	innerLeft.Add(
		&GridItem{
			p:       NewBlackBox(),
			row:     1,
			col:     0,
			rowSpan: 1,
			colSpan: 1,
		})

	tableItem :=
		&GridItem{
			p:       MakeTable(),
			row:     1,
			col:     1,
			rowSpan: 1,
			colSpan: 1,
			focus:   true,
		}

	innerLeft.Add(tableItem)

	innerGrid.Add(
		&GridItem{
			p:       innerLeft,
			row:     0, // row
			col:     0, // column
			rowSpan: 1, // rowSpan
			colSpan: 1, // colSpan
			focus:   true,
		})

	Grid.Add(
		&GridItem{
			p:       innerGrid, // MakeInnerLayout(),
			row:     1,         // row
			col:     1,         // column
			rowSpan: 1,         // rowSpan
			colSpan: 1,         // colSpan
			focus:   true,
		})

	return Grid.Grid, tableItem.p
}

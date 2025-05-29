package ui

import "github.com/rivo/tview"

type GridItem struct {
	p       tview.Primitive
	row     int
	col     int
	rowSpan int
	colSpan int
	focus   bool
	items   []*GridItem
}

type EasyGrid struct {
	*tview.Grid
}

func (g *EasyGrid) Add(gi *GridItem) {
	g.AddItem(gi.p, gi.row, gi.col, gi.rowSpan, gi.colSpan, 0, 0, gi.focus)
}

func NewEasyGrid() *EasyGrid {
	return &EasyGrid{
		tview.NewGrid(),
	}
}

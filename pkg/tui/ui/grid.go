package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Grid struct {
	*tview.Grid
	*TreeNode
	Rows            []int
	Columns         []int
	Title           string
	BorderColor     tcell.Color
	BackgroundColor tcell.Color
}

func (g *Grid) init() *GridPosition {
	g.SetColumns(g.Columns...)
	g.SetRows(g.Rows...)
	g.SetBackgroundColor(g.BackgroundColor)
	g.SetTitle(g.Title)
	if g.BorderColor != 0 {
		g.SetBorder(true)
		g.SetBorderColor(g.BorderColor)
	}
	for _, node := range g.children {
		pos := node.init()
		g.AddItem(node, pos.Row, pos.Column, pos.RowSpan, pos.ColSpan, pos.MinGridHeight, pos.MinGridWidth, pos.Focus)
	}
	return g.Pos
}

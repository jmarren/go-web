package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Box struct {
	*tview.Box
	*TreeNode
	BackgroundColor tcell.Color
	BorderColor     tcell.Color
}

func (b *Box) init() *GridPosition {
	if b.BackgroundColor != 0 {
		b.SetBackgroundColor(b.BackgroundColor)
	}
	if b.BorderColor != 0 {
		b.SetBorder(true)
		b.SetBorderColor(b.BorderColor)
	}
	return b.Pos
}

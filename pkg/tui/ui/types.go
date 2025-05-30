package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GridPosition struct {
	Row, Column, RowSpan, ColSpan, MinGridHeight, MinGridWidth int
	Focus                                                      bool
}

type EasyPrimitive interface {
	tview.Primitive
	ITreeNode
}

type InputMiddleware func(event *tcell.EventKey) *tcell.EventKey

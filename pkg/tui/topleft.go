package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewBlackBox() *tview.Box {
	return tview.NewBox().SetBackgroundColor(tcell.ColorBlack)

}

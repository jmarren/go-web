package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// var Box *tview.Box

func MakeBox() *tview.Box {
	return tview.NewBox().SetBorder(true).SetBorderColor(tcell.ColorWhite).SetTitle(" Go Web ")
}

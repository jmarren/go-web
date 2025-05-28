package tui

import (
	// "fmt"
	//
	// "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var Tui *tview.Application

func init() {
	layout, _ := MakeLayout()
	Tui = tview.NewApplication().SetRoot(layout, true).SetFocus(layout)

	Tui.SetFocus(layout)
}

package ui

import "github.com/gdamore/tcell/v2"

func ChainInputCapture(captures []func(event *tcell.EventKey) *tcell.EventKey) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		for _, capt := range captures {
			event = capt(event)
		}
		return event
	}
}

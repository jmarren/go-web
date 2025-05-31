package ui

import (
	"io"

	"github.com/gdamore/tcell/v2"
	"github.com/jmarren/go-web/pkg/utils"
	"github.com/rivo/tview"
)

type TextArea struct {
	*tview.TextArea
	*TreeNode
	BackgroundColor tcell.Color
	BorderColor     tcell.Color
	Text            []byte
}

func (t *TextArea) init() *GridPosition {
	t.SetBackgroundColor(t.BackgroundColor)
	if t.BorderColor != 0 {
		t.SetBorder(true)
		t.SetBorderColor(t.BorderColor)
	}
	return t.Pos
}

// drawFunc is used to update the screen after each write
func (t *TextArea) Writer(app *tview.Application) io.Writer {
	return utils.WriterFrom(func(p []byte) {
		app.QueueUpdateDraw(func() {
			t.AppendText(p)
		})
	})
}

func (t *TextArea) AppendText(p []byte) {
	// append p to CurrentText
	t.Text = append(t.Text, p...)

	// if CurrentText is longer than 1000,
	// take only the last 1000 bytes
	length := len(t.Text)
	if length > 1000 {
		t.Text = t.Text[length-1000:]
	}
	length = len(t.Text)

	if length > 4 {
		if isPrompt(t.Text[length-3:]) {
			t.Text = append(t.Text, '\n')
		}
	}

	// set textarea
	t.SetText(string(t.Text), true)

}

func isPrompt(b []byte) bool {
	if len(b) < 4 {
		return false
	}
	if b[0] == ':' && b[1] == '/' && b[2] == '$' {
		return true
	}
	return false
}

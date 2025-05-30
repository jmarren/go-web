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

func (t *TextArea) Writer() io.Writer {
	return utils.WriterFrom(t.AppendText)
}

func (t *TextArea) AppendText(p []byte) {
	// append p to CurrentText
	t.Text = append(t.Text, p...)

	// if CurrentText is longer than 500,
	// take only the last 500 bytes
	length := len(t.Text)
	if length > 500 {
		t.Text = t.Text[length-500:]
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

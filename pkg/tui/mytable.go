package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MyTable struct {
	*tview.Table
	data     [][]string
	selected *Cell
}

func (t *MyTable) HandleKey(event *tcell.EventKey) {
	t.ResetSelectedStyle()
	switch event.Key() {
	case tcell.KeyLeft:
		t.MoveLeft()
	case tcell.KeyRight:
		t.MoveRight()
	case tcell.KeyUp:
		t.MoveUp()
	case tcell.KeyDown:
		t.MoveDown()
	}
	t.UpdateSelectedStyle()
}

func (t *MyTable) UpdateSelectedStyle() {
	row := t.GetCurrRow()
	t.UpdateCells(row, func(cell *tview.TableCell) {
		cell.SetBackgroundColor(tcell.ColorYellow)
	})
}

func (t *MyTable) GetCurrRow() []*tview.TableCell {
	var cells []*tview.TableCell
	for i := range len(t.data[0]) {
		cells = append(cells, t.GetCell(t.selected.row, i))
	}
	return cells
}

func (t *MyTable) UpdateCells(cells []*tview.TableCell, callback func(cell *tview.TableCell)) {
	for _, cell := range cells {
		callback(cell)
	}
}

func (t *MyTable) ResetSelectedStyle() {
	row := t.GetCurrRow()
	t.UpdateCells(row, func(cell *tview.TableCell) {
		cell.SetBackgroundColor(tcell.ColorBlack)
	})
}

func (t *MyTable) GetCurr() *tview.TableCell {
	return t.GetCell(t.selected.row, t.selected.col)
}

func (t *MyTable) MoveLeft() {
	if t.selected.col > 1 {
		t.selected.MoveLeft()
	}
}

func (t *MyTable) MoveRight() {
	if t.selected.col < len(t.data[0])-1 {
		t.selected.MoveRight()
	}
}

func (t *MyTable) MoveUp() {
	if t.selected.row > 1 {
		t.selected.MoveUp()
	}
}

func (t *MyTable) MoveDown() {
	if t.selected.row < len(t.data)-1 {
		t.selected.MoveDown()
	}
}

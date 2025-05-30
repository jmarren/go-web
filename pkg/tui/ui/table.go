package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Table struct {
	*tview.Table
	*TreeNode
	BackgroundColor tcell.Color
	BorderColor     tcell.Color
	Title           string
	BorderPadding   []int
	Data            [][]string
	selected        *Cell
	middlewares     []func(event *tcell.EventKey) *tcell.EventKey
}

func (t *Table) init() *GridPosition {
	t.SetBackgroundColor(t.BackgroundColor)
	if t.BorderColor != 0 {
		t.SetBorder(true)
		t.SetBorderColor(t.BorderColor)
	}
	t.middlewares = []func(event *tcell.EventKey) *tcell.EventKey{t.baseInputCapture}

	t.SetTitle(t.Title)
	t.initCells()
	return t.Pos
}

func (t *Table) initCells() {
	for i, row := range t.Data {
		for column := range row {
			color := tcell.ColorWhite
			align := tview.AlignCenter

			if i == 0 {
				color = tcell.ColorYellow
			} else if column == 0 {
				align = tview.AlignLeft
				color = tcell.ColorDarkCyan
			}

			if row[column] == "online" {
				color = tcell.ColorGreen
			} else if row[column] == "offline" {
				color = tcell.ColorRed
			}

			cell := &tview.TableCell{
				Text:            row[column],
				Color:           color,
				Align:           align,
				BackgroundColor: tcell.ColorBlack,
			}

			t.SetCell(i, column, cell)
		}
	}
}

func (t *Table) updateInputCapture() {
	t.SetInputCapture(ChainInputCapture(t.middlewares))
}

func (t *Table) ShiftMiddleware(newMiddleware func(event *tcell.EventKey) *tcell.EventKey) {
	middlewares := []func(event *tcell.EventKey) *tcell.EventKey{newMiddleware}
	t.middlewares = append(middlewares, t.middlewares...)
	t.updateInputCapture()
}

func (t *Table) baseInputCapture(event *tcell.EventKey) *tcell.EventKey {
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
	return event
}

func (t *Table) UpdateSelectedStyle() {
	row := t.GetCurrRow()
	t.UpdateCells(row, func(cell *tview.TableCell) {
		cell.SetBackgroundColor(tcell.ColorYellow)
	})
}

func (t *Table) GetCurrRow() []*tview.TableCell {
	var cells []*tview.TableCell
	for i := range len(t.Data[0]) {
		cells = append(cells, t.GetCell(t.selected.row, i))
	}
	return cells
}

func (t *Table) UpdateCells(cells []*tview.TableCell, callback func(cell *tview.TableCell)) {
	for _, cell := range cells {
		callback(cell)
	}
}

func (t *Table) ResetSelectedStyle() {
	row := t.GetCurrRow()
	t.UpdateCells(row, func(cell *tview.TableCell) {
		cell.SetBackgroundColor(tcell.ColorBlack)
	})
}

func (t *Table) GetCurr() *tview.TableCell {
	return t.GetCell(t.selected.row, t.selected.col)
}

func (t *Table) MoveLeft() {
	if t.selected.col > 1 {
		t.selected.MoveLeft()
	}
}

func (t *Table) MoveRight() {
	if t.selected.col < len(t.Data[0])-1 {
		t.selected.MoveRight()
	}
}

func (t *Table) MoveUp() {
	if t.selected.row > 1 {
		t.selected.MoveUp()
	}
}

func (t *Table) MoveDown() {
	if t.selected.row < len(t.Data)-1 {
		t.selected.MoveDown()
	}
}

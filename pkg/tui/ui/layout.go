package ui

import (
	"github.com/jmarren/go-web/pkg/tui/config"
	"github.com/rivo/tview"
)

type Layout struct {
	*EasyGrid
	cfg *config.Grid
}

func NewLayout() *Layout {
	l := &Layout{
		EasyGrid: NewEasyGrid(),
		cfg:      config.Config.Ui.Layout.Grid,
	}
	l.SetColumns(l.cfg.Columns...)
	l.SetRows(l.cfg.Rows...)
	l.SetInner(NewInnerGrid())
	return l
}

func (l *Layout) SetInner(p tview.Primitive) {
	l.Add(&GridItem{
		p:       p, // MakeInnerLayout(),
		row:     1, // row
		col:     1, // column
		rowSpan: 1, // rowSpan
		colSpan: 1, // colSpan
		focus:   true,
	})
}

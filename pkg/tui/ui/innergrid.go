package ui

import "github.com/jmarren/go-web/pkg/tui/config"

type InnerGrid struct {
	*EasyGrid
}

func NewInnerGrid() *InnerGrid {
	cfg := config.Config.Ui.InnerGrid.Grid
	i := &InnerGrid{
		EasyGrid: NewEasyGrid(),
	}
	i.SetRows(cfg.Rows...)
	i.SetColumns(cfg.Columns...)
	i.SetTitle(cfg.Title)
	if cfg.BorderColor != 0 {
		i.SetBorderColor(cfg.BorderColor)
	}

	return i
}

// func (a *App) CreateInnerGrid() {
//
// 	InnerGrid := a.InnerGrid
// 	InnerGrid.SetRows()
// 	InnerGrid.SetColumns(-100, -100)
// 	InnerGrid.SetTitle(" Go Web! ")
// 	InnerGrid.SetBorder(true).SetBorderColor(tcell.ColorWhite)
//
// 	InnerGrid.Add(
// 		&GridItem{
// 			p:       a.InnerLeft,
// 			row:     0, // row
// 			col:     0, // column
// 			rowSpan: 1, // rowSpan
// 			colSpan: 1, // colSpan
// 			focus:   true,
// 		})
// 	a.InnerGrid = InnerGrid
// }

package ui

import "github.com/jmarren/go-web/pkg/tui/config"

type InnerLeft struct {
	*EasyGrid
}

func NewInnerLeft() *InnerLeft {
	i := &InnerLeft{
		EasyGrid: NewEasyGrid(),
	}
	cfg := config.Config.Ui.Layout.Inner.Left.Grid
	i.SetRows(cfg.Rows...)
	i.SetColumns(cfg.Columns...)
	i.SetBackgroundColor(cfg.BackgroundColor)
	return i
}

// func (a *App) CreateInnerLeft() {
// 	InnerLeft := a.InnerLeft
// 	InnerLeft.SetRows(-1, -1)
// 	InnerLeft.SetColumns(-1, -50, -1)
// 	InnerLeft.SetBackgroundColor(tcell.ColorBlack)
//
// 	InnerLeft.Add(
// 		&GridItem{
// 			p:       NewBlackBox(),
// 			row:     0,
// 			col:     0,
// 			rowSpan: 1,
// 			colSpan: 1,
// 		})
//
// 	// textArea := tview.NewTextArea()
// 	//
// 	// textArea.SetText(a.SshWrites, true)
// 	//
// 	InnerLeft.Add(
// 		&GridItem{
// 			p:       a.SshText,
// 			row:     0,
// 			col:     1,
// 			rowSpan: 1,
// 			colSpan: 1,
// 		})
//
// 	InnerLeft.Add(
// 		&GridItem{
// 			p:       NewBlackBox(),
// 			row:     1,
// 			col:     0,
// 			rowSpan: 1,
// 			colSpan: 1,
// 		})
//
// 	tableItem :=
// 		&GridItem{
// 			p:       MakeTable(),
// 			row:     1,
// 			col:     1,
// 			rowSpan: 1,
// 			colSpan: 1,
// 			focus:   true,
// 		}
//
// 	InnerLeft.Add(tableItem)
// 	a.InnerLeft = InnerLeft
// }

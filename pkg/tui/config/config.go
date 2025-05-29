package config

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Test() *tview.Application {
	app := tview.NewApplication()

	layout := Config.Ui
	layout.init()

	app.SetRoot(layout, true).SetFocus(layout)

	return app
}

type GridPosition struct {
	Row, Column, RowSpan, ColSpan, MinGridHeight, MinGridWidth int
	Focus                                                      bool
}

// func (g *Grid) AddItem(p Primitive, row, column, rowSpan, colSpan, minGridHeight, minGridWidth int, focus bool) *Grid {

type EasyPrimitive interface {
	tview.Primitive
	init() *GridPosition
}

type ConfigType struct {
	Ui *UiConfig
}

type UiConfig struct {
	*Grid
}

type Grid struct {
	*tview.Grid
	Name            string
	Rows            []int
	Columns         []int
	Title           string
	BorderColor     tcell.Color
	BackgroundColor tcell.Color
	Pos             *GridPosition
	Items           []EasyPrimitive
}

type LayoutConfig struct {
	Grid *Grid
}

func (g *Grid) init() *GridPosition {
	g.SetColumns(g.Columns...)
	g.SetRows(g.Rows...)
	g.SetBackgroundColor(g.BackgroundColor)
	g.SetTitle(g.Title)
	if g.BorderColor != 0 {
		g.SetBorder(true)
		g.SetBorderColor(g.BorderColor)
	}
	for _, node := range g.Items {
		pos := node.init()
		g.AddItem(node, pos.Row, pos.Column, pos.RowSpan, pos.ColSpan, pos.MinGridHeight, pos.MinGridWidth, pos.Focus)
	}
	return g.Pos
}

type Box struct {
	*tview.Box
	BackgroundColor tcell.Color
	Pos             *GridPosition
	BorderColor     tcell.Color
}

func (b *Box) init() *GridPosition {
	b.SetBackgroundColor(b.BackgroundColor)
	if b.BorderColor != 0 {
		b.SetBorder(true)
		b.SetBorderColor(b.BorderColor)
	}
	return b.Pos
}

var Config = &ConfigType{
	Ui: &UiConfig{
		Grid: &Grid{
			Grid:    tview.NewGrid(),
			Name:    "Layout",
			Rows:    []int{-1, -10, -1},
			Columns: []int{-1, -10, -1},
			Pos: &GridPosition{
				RowSpan: -1,
				ColSpan: -1,
			},
			Items: []EasyPrimitive{
				&Grid{
					Grid:        tview.NewGrid(),
					Name:        "InnerLayout",
					Rows:        []int{},
					Columns:     []int{-1, -10, -1},
					Title:       " Go Web! ",
					BorderColor: tcell.ColorWhite,
					Pos: &GridPosition{
						Row:     1,
						Column:  1,
						RowSpan: 1,
						ColSpan: 1,
						Focus:   true,
					},
					Items: []EasyPrimitive{
						&Grid{
							Grid:            tview.NewGrid(),
							Name:            "InnerLeft",
							Rows:            []int{-1, -1},
							Columns:         []int{-1, -50, -1},
							Title:           "InnerLeft",
							BackgroundColor: tcell.ColorBlack,
							Pos: &GridPosition{
								Row:           1,
								Column:        1,
								RowSpan:       1,
								ColSpan:       1,
								MinGridWidth:  0,
								MinGridHeight: 0,
								Focus:         true,
							},
							Items: []EasyPrimitive{
								&Box{
									Box:             tview.NewBox(),
									BackgroundColor: tcell.ColorBlue,
									BorderColor:     tcell.ColorYellow,
									Pos: &GridPosition{
										Row:           1,
										Column:        1,
										RowSpan:       1,
										ColSpan:       1,
										MinGridWidth:  0,
										MinGridHeight: 0,
										Focus:         true,
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

// Inner: &InnerGridConfig{
// 	Grid: &Grid{
// 		Rows:        []int{},
// 		Columns:     []int{-1, -10, -1},
// 		Title:       " Go Web! ",
// 		BorderColor: tcell.ColorWhite,
// 	},
// 	Left: &InnerLeftConfig{
// 		Grid: &Grid{
// 			Rows:            []int{-1, -1},
// 			Columns:         []int{-1, -50, -1},
// 			BackgroundColor: tcell.ColorBlack,
// 		},
// 	},
// },

// type ConfigT struct {
// 	Ui
// }

type GridT struct {
	Rows            []int
	Columns         []int
	Title           string
	BorderColor     tcell.Color
	BackgroundColor tcell.Color
	items           []tview.Primitive
}

type Node struct {
	p    tview.Primitive
	attr any

	// grid tview.Grid
}

// InnerLeft.SetRows(-1, -1)
// InnerLeft.SetColumns(-1, -50, -1)
// InnerLeft.SetBackgroundColor(tcell.ColorBlack)
//
// var Config = struct {
// 	Ui struct {
// 		Layout struct {
// 			rows    []int
// 			columns []int
// 		}
// 	}
// }{
// 	Ui: struct {
// 		Layout: struct{
// 			rows:    []int{-1, -10, -1},
// 			columns: []int{-1, -10, -1},
// 		},
// 	},
// }

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

type EasyPrimitive interface {
	tview.Primitive
	init() *GridPosition
	GetName() string
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

func (g *Grid) GetName() string {
	return g.Name
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
	Name            string
	BackgroundColor tcell.Color
	Pos             *GridPosition
	BorderColor     tcell.Color
}

func (b *Box) GetName() string {
	return b.Name
}

type TextArea struct {
	*tview.TextArea
	Name            string
	BackgroundColor tcell.Color
	Pos             *GridPosition
	BorderColor     tcell.Color
}

func (t *TextArea) GetName() string {
	return t.Name
}

func (t *TextArea) init() *GridPosition {
	t.SetBackgroundColor(t.BackgroundColor)
	if t.BorderColor != 0 {
		t.SetBorder(true)
		t.SetBorderColor(t.BorderColor)
	}
	return t.Pos
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
					Grid:            tview.NewGrid(),
					Name:            "InnerLayout",
					Rows:            []int{},
					Columns:         []int{-1, -1},
					Title:           " Go Web! ",
					BorderColor:     tcell.ColorWhite,
					BackgroundColor: tcell.ColorBlack,
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
								Row:           0,
								Column:        0,
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
								&Box{
									Box:             tview.NewBox(),
									BackgroundColor: tcell.ColorBlack,
									Pos: &GridPosition{
										Row:     0,
										Column:  0,
										RowSpan: 2,
										ColSpan: 1,
									},
								},
								&TextArea{
									TextArea:        tview.NewTextArea(),
									BackgroundColor: tcell.ColorAliceBlue,
									BorderColor:     tcell.ColorRed,
									Pos: &GridPosition{
										Row:     0,
										Column:  1,
										RowSpan: 1,
										ColSpan: 1,
									},
								},
								&Table{
									Table:           tview.NewTable(),
									BackgroundColor: tcell.ColorBlack,
									Title:           " instances ",
									Data: [][]string{
										{"instance", "status", "uptime", "playbooks"},
										{"devdb", "online", "12m 15s", "db"},
										{"devapp", "online", "10m 05s", "app"},
										{"app", "offline", "--", "--"},
										{"db", "offline", "--", "--"},
									},
									Pos: &GridPosition{
										Row:     1,
										Column:  1,
										RowSpan: 1,
										ColSpan: 1,
										Focus:   true,
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

type Cell struct {
	row int
	col int
}

func (c *Cell) MoveLeft() {
	c.col--
}

func (c *Cell) MoveRight() {
	c.col++
}

func (c *Cell) MoveDown() {
	c.row++
}

func (c *Cell) MoveUp() {
	c.row--
}

type Table struct {
	*tview.Table
	Name            string
	BackgroundColor tcell.Color
	BorderColor     tcell.Color
	Title           string
	BorderPadding   []int
	Data            [][]string
	Selected        *Cell
	Pos             *GridPosition
}

func (t *Table) GetName() string {
	return t.Name
}

func (t *Table) init() *GridPosition {
	t.SetBackgroundColor(t.BackgroundColor)
	if t.BorderColor != 0 {
		t.SetBorder(true)
		t.SetBorderColor(t.BorderColor)
	}
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

var rows = [][]string{
	{"instance", "status", "uptime", "playbooks"},
	{"devdb", "online", "12m 15s", "db"},
	{"devapp", "online", "10m 05s", "app"},
	{"app", "offline", "--", "--"},
	{"db", "offline", "--", "--"},
}

// func MakeTable() *tview.Table {
// 	Table := tview.NewTable()
// 	tbl := &MyTable{
// 		Table: Table,
// 		data:  rows,
// 		selected: &Cell{
// 			row: 1,
// 			col: 1,
// 		},
// 	}
//
// 	Table.SetBackgroundColor(tcell.ColorBlack)
// 	Table.SetBorder(true).SetBorderColor(tcell.ColorWhite).SetTitle(" Instances ")
// 	Table.SetBorderPadding(1, 1, 2, 2)
// 	Table.SetSelectionChangedFunc(func(row, column int) {
// 		fmt.Printf("row %d col %d selected\n", row, column)
// 	})
//
// 	Table.SetSelectedFunc(func(row, column int) {
// 		fmt.Printf("row %d col %d selected\n", row, column)
// 	})
// 	for i, row := range rows {
// 		for column := range row {
// 			color := tcell.ColorWhite
// 			align := tview.AlignCenter
//
// 			if i == 0 {
// 				color = tcell.ColorYellow
// 			} else if column == 0 {
// 				align = tview.AlignLeft
// 				color = tcell.ColorDarkCyan
// 			}
//
// 			if row[column] == "online" {
// 				color = tcell.ColorGreen
// 			} else if row[column] == "offline" {
// 				color = tcell.ColorRed
// 			}
//
// 			cell := &tview.TableCell{
// 				Text:            row[column],
// 				Color:           color,
// 				Align:           align,
// 				BackgroundColor: tcell.ColorBlack,
// 			}
//
// 			Table.SetCell(i, column, cell)
// 		}
// 	}
//
// 	Table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
// 		tbl.HandleKey(event)
// 		return event
// 	})
//
// 	return Table
// }

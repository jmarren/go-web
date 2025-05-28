package tui

import (
	// "fmt"
	// "os/exec"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GridItem struct {
	p       tview.Primitive
	row     int
	col     int
	rowSpan int
	colSpan int
	focus   bool
}

type MyGrid struct {
	*tview.Grid
}

func (g *MyGrid) Add(gi *GridItem) {
	g.AddItem(gi.p, gi.row, gi.col, gi.rowSpan, gi.colSpan, 0, 0, gi.focus)
}

func NewMyGrid() *MyGrid {
	return &MyGrid{
		tview.NewGrid(),
	}
}

// func NewApp() *App {
// 	return &App{
// 		MyGrid:    newMyGrid(),
// 		InnerGrid: newMyGrid(),
// 		InnerLeft: newMyGrid(),
// 	}
// }

func (a *App) CreateLayout() {
	rows := []int{-1, -10, -1}
	columns := []int{-1, -10, -1}

	Layout := a.Layout

	Layout.SetRows(rows...)
	Layout.SetColumns(columns...)

	Layout.Add(
		&GridItem{
			p:       a.InnerGrid, // MakeInnerLayout(),
			row:     1,           // row
			col:     1,           // column
			rowSpan: 1,           // rowSpan
			colSpan: 1,           // colSpan
			focus:   true,
		})
	a.Layout = Layout
}

func (a *App) CreateInnerGrid() {

	InnerGrid := a.InnerGrid
	InnerGrid.SetRows()
	InnerGrid.SetColumns(-100, -100)
	InnerGrid.SetTitle(" Go Web! ")
	InnerGrid.SetBorder(true).SetBorderColor(tcell.ColorWhite)

	InnerGrid.Add(
		&GridItem{
			p:       a.InnerLeft,
			row:     0, // row
			col:     0, // column
			rowSpan: 1, // rowSpan
			colSpan: 1, // colSpan
			focus:   true,
		})
	a.InnerGrid = InnerGrid
}

func (a *App) CreateInnerLeft() {
	InnerLeft := a.InnerLeft
	InnerLeft.SetRows(-1, -1)
	InnerLeft.SetColumns(-1, -50, -1)
	InnerLeft.SetBackgroundColor(tcell.ColorBlack)

	InnerLeft.Add(
		&GridItem{
			p:       NewBlackBox(),
			row:     0,
			col:     0,
			rowSpan: 1,
			colSpan: 1,
		})

	// textArea := tview.NewTextArea()
	//
	// textArea.SetText(a.SshWrites, true)
	//
	InnerLeft.Add(
		&GridItem{
			p:       a.SshText,
			row:     0,
			col:     1,
			rowSpan: 1,
			colSpan: 1,
		})

	InnerLeft.Add(
		&GridItem{
			p:       NewBlackBox(),
			row:     1,
			col:     0,
			rowSpan: 1,
			colSpan: 1,
		})

	tableItem :=
		&GridItem{
			p:       MakeTable(),
			row:     1,
			col:     1,
			rowSpan: 1,
			colSpan: 1,
			focus:   true,
		}

	InnerLeft.Add(tableItem)
	a.InnerLeft = InnerLeft
}

func (a *App) CreateSshTextArea() {
	// tv := tview.NewTextView()
	ta := tview.NewTextArea()
	ta.SetText("laldksfji", true)
	a.SshText = ta
}

// func MakeLayout() (*tview.Grid, tview.Primitive) {
// 	app := NewApp()
//
// 	rows := []int{-1, -10, -1}
// 	columns := []int{-1, -10, -1}
//
// 	app.SetRows(rows...)
// 	app.SetColumns(columns...)
//
// 	// InnerGrid := newMyGrid()
// 	InnerGrid := app.InnerGrid
// 	InnerGrid.SetRows()
// 	InnerGrid.SetColumns(-100, -100)
// 	InnerGrid.SetTitle(" Go Web! ")
// 	InnerGrid.SetBorder(true).SetBorderColor(tcell.ColorWhite)
//
// 	InnerLeft := app.InnerLeft
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
// 	textArea := tview.NewTextArea()
//
// 	// textArea.SetText(app.sshWrites, true)
//
// 	InnerLeft.Add(
// 		&GridItem{
// 			p:       textArea,
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
//
// 	InnerGrid.Add(
// 		&GridItem{
// 			p:       InnerLeft,
// 			row:     0, // row
// 			col:     0, // column
// 			rowSpan: 1, // rowSpan
// 			colSpan: 1, // colSpan
// 			focus:   true,
// 		})
//
// 	app.Add(
// 		&GridItem{
// 			p:       InnerGrid, // MakeInnerLayout(),
// 			row:     1,         // row
// 			col:     1,         // column
// 			rowSpan: 1,         // rowSpan
// 			colSpan: 1,         // colSpan
// 			focus:   true,
// 		})
//
// 	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
// 		// if event.Key() == tcell.KeyEnter {
// 		// 	app.handleEnter()
// 		// }
// 		return event
// 	})
//
// 	return app.Grid, tableItem.p
// }
//
// func (a *App) Write(p []byte) (int, error) {
// 	// fmt.Printf("%s", string(p))
// 	// a.sshWrites = a.sshWrites + string(p)
// 	textArea := tview.NewTextArea()
//
// 	textArea.SetText("hi", true)
//
// 	a.InnerLeft.Add(
// 		&GridItem{
// 			p:       textArea,
// 			row:     0,
// 			col:     1,
// 			rowSpan: 1,
// 			colSpan: 1,
// 		})
//
// 	return len(p), nil
// }
//
// type myWriter struct{}
//
// func (w *myWriter) Write(p []byte) (int, error) {
// 	// fmt.Printf("%s", string(p))
// 	// a.sshWrites = a.sshWrites + string(p)
// 	return len(p) + 1, nil
// }
//
// func (a *App) handleEnter() {
//
// 	// go func() {
// 	// 	Tui.Suspend(func() {
// 	// 		cmd := exec.Command("myapp", "p")
// 	// 		_, err := cmd.Output()
// 	//
// 	// 		if err != nil {
// 	// 			fmt.Println(err.Error())
// 	// 		}
// 	// 	})
// 	//
// 	// 	// Print the output
// 	// 	// fmt.Println(string(stdout))
// 	// }()
// }

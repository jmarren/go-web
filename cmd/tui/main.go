package main

import (
	"fmt"
	// "math"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TableData struct {
	tview.TableContentReadOnly
}

func (d *TableData) GetCell(row, column int) *tview.TableCell {
	// letters := [...]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 'A' + byte(row%26)} // log(math.MaxInt64) / log(26) ~= 14

	letters := []byte{'H', 'i', '!'}

	// start := len(letters) - 1
	// row /= 26
	// for row > 0 {
	// 	start--
	// 	row--
	// 	letters[start] = 'A' + byte(row%26)
	// 	row /= 26
	// }
	return tview.NewTableCell(fmt.Sprintf("[red]%s[green]%d", letters, column))
}

func (d *TableData) GetRowCount() int {
	return 2
}

func (d *TableData) GetColumnCount() int {
	return 2
}

func main() {

	table := tview.NewTable().SetFixed(1, 1)

	for row := 0; row < 100; row++ {
		for column := 0; column < 7; column++ {
			color := tcell.ColorWhite
			if row == 0 {
				color = tcell.ColorYellow
			} else if column == 0 {
				color = tcell.ColorDarkCyan
			}

			align := tview.AlignLeft

			if row == 0 {
				align = tview.AlignCenter
			} else if column == 0 || column >= 4 {
				align = tview.AlignRight
			}

			text := "..."

			if row == 0 {
				text = "header"
			} else if column == 0 || column >= 4 {
				text = "row"
			}

			table.SetCell(row, column, &tview.TableCell{
				Text:  text,
				Color: color,
				Align: align,
			})
		}
	}

	// layout := tview.NewGrid().SetBackgroundColor(tcell.ColorGreen)
	//
	// 	layout. := layout
	// // modal := tview.NewModal()
	// //
	// // modal.SetText("hi")
	//
	app := tview.NewApplication().SetRoot(table, true)

	if err := app.Run(); err != nil {
		panic(err)
	}

	// data := &TableData{}
	// table := tview.NewTable().
	// 	SetBorders(true).
	// 	SetSelectable(true, true).
	// 	SetContent(data)
	// if err := tview.NewApplication().SetRoot(table, true).EnableMouse(true).Run(); err != nil {
	// 	panic(err)
	// }
}

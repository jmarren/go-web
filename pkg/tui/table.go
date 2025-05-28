package tui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var rows = [][]string{
	{"instance", "status", "uptime", "playbooks"},
	{"devdb", "online", "12m 15s", "db"},
	{"devapp", "online", "10m 05s", "app"},
	{"app", "offline", "--", "--"},
	{"db", "offline", "--", "--"},
}

func MakeTable() *tview.Table {
	Table := tview.NewTable()
	tbl := &MyTable{
		Table: Table,
		data:  rows,
		selected: &Cell{
			row: 1,
			col: 1,
		},
	}

	Table.SetBackgroundColor(tcell.ColorBlack)
	Table.SetBorder(true).SetBorderColor(tcell.ColorWhite).SetTitle(" Instances ")
	Table.SetBorderPadding(1, 1, 2, 2)
	Table.SetSelectionChangedFunc(func(row, column int) {
		fmt.Printf("row %d col %d selected\n", row, column)
	})

	Table.SetSelectedFunc(func(row, column int) {
		fmt.Printf("row %d col %d selected\n", row, column)
	})
	for i, row := range rows {
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

			Table.SetCell(i, column, cell)
		}
	}

	Table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		tbl.HandleKey(event)
		return event
	})

	return Table
}

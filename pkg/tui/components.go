package tui

import "github.com/jmarren/go-web/pkg/tui/ui"

func (t *Tui) terminal() *ui.TextArea {
	terminal, err := ui.GetById[*ui.TextArea](t.ui.Grid, ui.TerminalArea)
	if err != nil {
		t.error(err)
	}
	return *terminal
}

func (t *Tui) instanceTable() *ui.Table {
	table, err := ui.GetById[*ui.Table](t.ui.Grid, ui.InstanceTable)
	if err != nil {
		t.error(err)
	}
	return *table
}

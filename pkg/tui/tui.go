package tui

import (
	"errors"

	"github.com/gdamore/tcell/v2"
	"github.com/jmarren/go-web/pkg/tui/logger"
	"github.com/jmarren/go-web/pkg/tui/myssh"
	"github.com/jmarren/go-web/pkg/tui/ui"
	"github.com/rivo/tview"
)

type Tui struct {
	ssh    *myssh.SshService
	ui     *ui.Ui
	logger *logger.Logger
	app    *tview.Application
}

func New() *Tui {

	t := &Tui{
		app: tview.NewApplication(),
	}

	logger, err := logger.New("app.log")
	if err != nil {
		t.error(err)
	}
	t.logger = logger

	t.ssh = myssh.New()

	t.InitUi()

	instanceTable := t.instanceTable()

	t.app.SetRoot(t.ui.Grid, true).SetFocus(instanceTable)
	return t
}

func (t *Tui) error(e ...error) {
	panic(errors.Join(e...))
}

func (t *Tui) activateTerminal() {
	err := t.ssh.Connect("devdb", t.WriteSsh)
	if err != nil {
		t.error(err)
	}
	terminal := t.terminal()
	t.app.SetFocus(terminal)
	terminal.SetInputCapture(t.captureTerminalInput)
}

func (t *Tui) captureTerminalInput(event *tcell.EventKey) *tcell.EventKey {
	if t.ssh.Active {
		t.ssh.PipeIn([]byte{byte(event.Rune())})
	}
	return event
}

func (t *Tui) captureInstanceTableInput(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyCtrlA && !t.ssh.Active {
		t.activateTerminal()
		return event
	}

	return event
}

func (t *Tui) InitUi() {
	t.ui = ui.New()
	t.instanceTable().ShiftMiddleware(t.captureInstanceTableInput)
}

// func addMiddleware(p ui.EasyPrimitive, middleware func(event *tcell.EventKey) *tcell.EventKey) {
// 	// capture := p.GetInpu
// 	handler := p.InputHandler()
//
// }

// func (t *Tui)

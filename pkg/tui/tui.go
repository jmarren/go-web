package tui

import (
	"errors"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/jmarren/go-web/pkg/tui/instance"
	"github.com/jmarren/go-web/pkg/tui/logger"
	"github.com/jmarren/go-web/pkg/tui/myssh"
	"github.com/jmarren/go-web/pkg/tui/ui"
	"github.com/rivo/tview"
	"golang.org/x/crypto/ssh"
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
	terminal := t.terminal()
	instance := t.instanceTable().SelectedInstance()

	err := t.ssh.Connect(instance.TermConfig, terminal.Writer())
	if err != nil {
		t.error(err)
	}
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
	instanceTable := t.instanceTable()
	instanceTable.ShiftMiddleware(t.captureInstanceTableInput)
	instanceTable.SetInstances(instances)
}

var instances = []*instance.Instance{
	&instance.Instance{
		Name:      "devdb",
		Online:    true,
		StartTime: time.Now().Add(-1*time.Hour - 3*time.Minute),
		TermConfig: &myssh.TermConfig{
			Addr:    "127.0.0.1:200",
			Network: "tcp",
			ClientCfg: &ssh.ClientConfig{
				User: "test",
				Auth: []ssh.AuthMethod{
					ssh.Password("test"),
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			},
		},
	},
}

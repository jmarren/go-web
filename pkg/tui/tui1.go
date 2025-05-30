package tui

/*
import (
	"errors"
	"io"

	"github.com/gdamore/tcell/v2"
	"github.com/jmarren/go-web/pkg/tui/logger"
	"github.com/jmarren/go-web/pkg/tui/myssh"
	"github.com/jmarren/go-web/pkg/tui/ui"
	"github.com/rivo/tview"
)

type App struct {
	*tview.Application
	ui          *ui.Grid
	Layout      *MyGrid
	InnerGrid   *MyGrid
	InnerLeft   *MyGrid
	SshText     *tview.TextArea
	CurrentText []byte
	InPipe      io.WriteCloser
	logger      *logger.Logger
	sshService  *myssh.SshService
}

func Test() *tview.Application {
	app := tview.NewApplication()
	ui := ui.New()
	// app.SetInputCapture()
	app.SetRoot(ui.Grid, true).SetFocus(ui.Grid)
	return app
}

func Create() *App {
	app := &App{
		// Application: tview.NewApplication(),
		Layout:      NewMyGrid(),
		InnerGrid:   NewMyGrid(),
		InnerLeft:   NewMyGrid(),
		CurrentText: []byte{},
	}

	logger, err := logger.New("app.log")
	if err != nil {
		app.error(err)
	}
	app.logger = logger

	app.sshService = myssh.New(newSshWriter(app.AppendSshText))

	app.createUi()

	return app
}

func (a *App) captureInput(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyCtrlA && !a.sshService.Active {
		err := a.sshService.Connect("devdb")
		if err != nil {
			a.error(err)
		}
		return event
	}

	if a.sshService.Active {
		a.sshService.PipeIn([]byte{byte(event.Rune())})
	}

	return event
}

func (a *App) createUi() {
	a.CreateSshTextArea()
	a.CreateInnerLeft()
	a.CreateInnerGrid()
	a.CreateLayout()
	a.SetRoot(a.Layout, true).SetFocus(a.Layout)
	a.SetInputCapture(a.captureInput)
}

func (a *App) error(e ...error) {
	panic(errors.Join(e...))
}

type SshWriter struct {
	writeFunc func(p []byte)
}

func (s *SshWriter) Write(p []byte) (int, error) {
	s.writeFunc(p)
	return len(p), nil
}

func newSshWriter(writeFunc func(p []byte)) io.Writer {
	return &SshWriter{
		writeFunc: writeFunc,
	}
}

func (a *App) AppendSshText(p []byte) {

	// append p to CurrentText
	a.CurrentText = append(a.CurrentText, p...)

	// if CurrentText is longer than 500,
	// take only the last 500 bytes
	length := len(a.CurrentText)
	if length > 500 {
		a.CurrentText = a.CurrentText[length-500:]
	}
	// set textarea
	a.SshText.SetText(string(a.CurrentText), true)
	a.SetFocus(a.SshText)
}

func (a *App) Write(p []byte) (int, error) {
	a.AppendSshText(p)
	return len(p), nil
}

func (a *App) Read(p []byte) (n int, err error) {
	a.AppendSshText(p)
	return len(p), nil
}
*/

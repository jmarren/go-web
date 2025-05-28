package main

import (
	"github.com/jmarren/go-web/pkg/tui"
	"github.com/rivo/tview"
)

func main() {

	MyApp := &tui.App{
		Application: tview.NewApplication(),
		Layout:      tui.NewMyGrid(),
		InnerGrid:   tui.NewMyGrid(),
		InnerLeft:   tui.NewMyGrid(),
		SshWrites:   "",
	}
	app := tui.Create(MyApp)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

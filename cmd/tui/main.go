package main

import (
	"github.com/jmarren/go-web/pkg/tui"
)

func main() {

	app := tui.Test()
	if err := app.Run(); err != nil {
		panic(err)
	}
}

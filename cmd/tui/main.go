package main

import (
	"github.com/jmarren/go-web/pkg/tui"
)

func main() {
	if err := tui.Tui.Run(); err != nil {
		panic(err)
	}
}

package tui

import (
	"fmt"
	"os/exec"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// var Tui *tview.Application

type App struct {
	*tview.Application
	Layout    *MyGrid
	InnerGrid *MyGrid
	InnerLeft *MyGrid
	SshWrites string
	SshText   *tview.TextArea
}

type MyWriter struct{}

// Write(p []byte) (n int, err error)

func (w *MyWriter) Write(p []byte) (int, error) {
	fmt.Printf("p: %s\n", string(p))
	return len(p), nil
}

var MyApp *App

func Create(initial *App) *App {

	MyApp = initial

	MyApp.CreateSshTextArea()

	MyApp.CreateInnerLeft()

	MyApp.CreateInnerGrid()

	MyApp.CreateLayout()

	MyApp.SetRoot(MyApp.Layout, true).SetFocus(MyApp.Layout)

	MyApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			MyApp.Suspend(func() {
				go func() {
					cmd := exec.Command("myapp", "connect", "devdb")
					cmd.Stdout = MyApp
					cmd.Stderr = MyApp
					cmd.Stdin = MyApp
					err := cmd.Run()
					if err != nil {
						fmt.Printf("error: %s \n", err)
					}
				}()
			})
		} else {
			MyApp.Read([]byte{byte(event.Rune())})
		}
		return event
	})

	return MyApp

}

func (a *App) Write(p []byte) (int, error) {
	MyApp.SshWrites += string(p)
	a.SshText.SetText(MyApp.SshWrites, true)
	return len(p), nil
}

func (a *App) Read(p []byte) (n int, err error) {
	return 0, nil
}

// Read(p []byte) (n int, err error)
// layout, _ := MakeLayout()

// Tui = tview.NewApplication().SetRoot(layout, true).SetFocus(layout)
//
// Tui.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
// 	if event.Key() == tcell.KeyEnter {
// 		Tui.Suspend(func() {
// 			cmd := exec.Command("myapp", "p")
// 			_, err := cmd.Output()
//
// 			if err != nil {
// 				fmt.Println(err.Error())
// 			}
// 		})
// 	}
// 	return event
// })
// MyApp.Layout.SetTitle("hi")
// MyApp.SetRoot(MyApp.Layout, true)
// textArea := tview.NewTextArea()
//
// textArea.SetText(MyApp.SshWrites, true)
//
// MyApp.InnerLeft.Add(
// 	&GridItem{
// 		p:       textArea,
// 		row:     0,
// 		col:     1,
// 		rowSpan: 1,
// 		colSpan: 1,
// 	})
//
// Create(MyApp)
// MyApp.CreateInnerLeft()
// MyApp.CreateInnerGrid()
// MyApp.CreateLayout()
/*
	var wg sync.WaitGroup
	//
	ch := make(chan string, 1)

	MyApp.Lock()
	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd := exec.Command("echo", "hi")
		cmd.Run()
		// bytes, err := cmd.Output()
		// MyApp.Stop()
		// if err != nil {
		// 	fmt.Printf("error: %s\n", err)
		// }
		// ch <- string(bytes)
		// Out, err := cmd.StdoutPipe()
		// fmt.Println(err)
		// cmd.Start()
		// outByte, _ := io.ReadAll(Out)
		// output := string(outByte)
		// fmt.Printf("output: %s \n", output)
		// ch <- string(output)
		// Out.Close()
		// err = cmd.Wait()
		// fmt.Println(err)
	}()
	wg.Wait()
	output := <-ch
	MyApp.SshWrites += output
	fmt.Printf("output: %s \n", output)
	MyApp.Unlock()
*/

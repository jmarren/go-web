package tui

import (
	"fmt"
	"io"
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
	SshOpen   bool
	InPipe    io.WriteCloser
	OutPipe   io.ReadCloser
	cmd       *exec.Cmd
}

// type MyWriter struct{}
//
// // Write(p []byte) (n int, err error)
//
// func (w *MyWriter) Write(p []byte) (int, error) {
// 	fmt.Printf("p: %s\n", string(p))
// 	return len(p), nil
// }

var MyApp *App

func Create(initial *App) *App {

	MyApp = initial
	MyApp.SshOpen = false
	MyApp.SshWrites = ""
	MyApp.cmd = exec.Command("myapp", "connect", "devdb")

	// // capture outpipe
	// outPipe, err := cmd.StdoutPipe()
	// if err != nil {
	// 	fmt.Printf("error: %s \n", err)
	// }
	// // capture in pipe
	// inPipe, err := cmd.StdinPipe()
	// if err != nil {
	// 	fmt.Printf("error: %s \n", err)
	// }
	//
	MyApp.CreateSshTextArea()

	MyApp.CreateInnerLeft()

	MyApp.CreateInnerGrid()

	MyApp.CreateLayout()

	MyApp.SetRoot(MyApp.Layout, true).SetFocus(MyApp.Layout)

	var sshopen = false

	MyApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter && !sshopen {
			sshopen = true
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
			return event
		}
		// if MyApp.SshOpen {
		// 	MyApp.Read([]byte{byte(event.Rune())})
		// }
		return event
	})

	return MyApp
}

func (a *App) Write(p []byte) (int, error) {
	a.SshWrites += string(p)
	a.SshText.SetText(a.SshWrites, true)
	a.SetFocus(a.SshText)
	return 0, nil
}

func (a *App) Read(p []byte) (n int, err error) {
	// a.SshWrites += string(p)
	// a.SshText.SetText(a.SshWrites, true)
	// a.SetFocus(a.SshText)
	return 0, nil
}
func (a *App) Close() error {

	return nil
}

// func (a *App) CaptureInput(event *tcell.EventKey) *tcell.EventKey {
//
// 	if event.Key() == tcell.KeyCtrlA {
// 		// if ssh is not open, open
// 		if !a.SshOpen {
// 			a.SshOpen = true
//
// 			a.Suspend(func() {
// 				var wg sync.WaitGroup
// 				wg.Add(1)
// 				go func() {
// 					defer wg.Done()
// 					err := a.cmd.Start()
// 					if err != nil {
// 						fmt.Println(err.Error())
// 						return
// 					}
// 					a.cmd.Stdout = a
// 					a.cmd.Stdin = a
// 					a.cmd.Stderr = a
// 					a.cmd.Wait()
// 				}()
// 				wg.Wait()
// 			})
// 		}
// 		return event
// 	}
//
// 	// if event.Key() == tcell.KeyEnter && a.SshOpen {
// 	// 	a.cmd.Wait()
// 	// 	_, err := io.ReadAll(a)
// 	// 	if err != nil {
// 	// 		fmt.Println(err.Error())
// 	// 	}
// 	// 	a.Close()
// 	// 	// a.Write(output)
// 	// 	// a.
// 	// }
//
// 	if a.SshOpen {
// 		// a.Write(event.Rune())
//
// 		if event.Key() == tcell.KeyEnter {
// 			fmt.Println(string(event.Rune()))
// 			// inpipe := a.cmd.Stdin
// 			_, err := a.Write([]byte{byte(event.Rune())})
// 			if err != nil {
// 				fmt.Println(err.Error())
// 			}
// 			a.cmd.Wait()
// 		}
//
// 		// fmt.Printf("wrote %d bytes\n", n)
// 		return event
// 	}
//
// 	return event
// }
//
/*

	if event.Key() == tcell.KeyEnter && !a.SshOpen {
		a.SshOpen = true
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
		return event
	}
	if a.SshOpen {
		MyApp.Read([]byte{byte(event.Rune())})
	}
	return event
}

/*

	if event.Key() == tcell.KeyCtrlA {
		// if ssh is not open, open
		if !a.SshOpen {
			a.SshOpen = true

			a.Suspend(func() {
				err := a.cmd.Start()
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				a.cmd.Stdout = a
				a.cmd.Stdin = a
				a.cmd.Wait()
				_, err = io.ReadAll(a)
				if err != nil {
					fmt.Println(err.Error())
				}
				// inpipe, err := a.cmd.StdinPipe()
				// if err != nil {
				// 	fmt.Println(err.Error())
				// }
				// a.InPipe = inpipe
				// io.ReadAll(a.OutPipe)
				// a.cmd.Stdout = a
				// a.cmd.Wait()
			})
		}
		return event
	}

	// if event.Key() == tcell.KeyEnter && a.SshOpen {
	// 	a.cmd.Wait()
	// 	_, err := io.ReadAll(a)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	a.Close()
	// 	// a.Write(output)
	// 	// a.
	// }

	if a.SshOpen {
		// inpipe := a.cmd.Stdin
		n, err := a.Read([]byte{byte(event.Rune())})
		// n, err := a.Write()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
		return event
	}

	return event
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
// output, err := io.ReadAll(outPipe)
// MyApp.Write(output)
// cmd.Stdout = MyApp
// cmd.Stdin = MyApp
// err = cmd.Run()
// if err != nil {
// 	fmt.Printf("error: %s \n", err)
// }
// }()

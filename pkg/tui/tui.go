package tui

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.org/x/crypto/ssh"
)

type App struct {
	*tview.Application
	Layout      *MyGrid
	InnerGrid   *MyGrid
	InnerLeft   *MyGrid
	SshText     *tview.TextArea
	CurrentText string
	SshOpen     bool
	InPipe      io.WriteCloser
	LogFile     string
	OutPipe     io.Reader
	session     *ssh.Session
	sshClient   *ssh.Client
}

func Create() *App {

	MyApp := &App{
		Application: tview.NewApplication(),
		Layout:      NewMyGrid(),
		InnerGrid:   NewMyGrid(),
		InnerLeft:   NewMyGrid(),
		LogFile:     "app.log",
		SshOpen:     false,
		CurrentText: "",
	}

	MyApp.InitFiles()

	MyApp.WriteLog("hi")

	MyApp.CreateSshTextArea()

	MyApp.CreateInnerLeft()

	MyApp.CreateInnerGrid()

	MyApp.CreateLayout()

	MyApp.SetRoot(MyApp.Layout, true).SetFocus(MyApp.Layout)

	// MyApp.Tick()
	MyApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlA && !MyApp.SshOpen {
			MyApp.SetClient()
			MyApp.SshOpen = true
			return event
		}
		MyApp.Write([]byte{byte(event.Rune())})
		if MyApp.SshOpen {
			MyApp.InPipe.Write([]byte{byte(event.Rune())})
		}

		return event
	})
	return MyApp
}

func (a *App) check(e error, s string) {
	if e != nil {
		a.WriteLog(fmt.Sprintf("error %s:  %s", s, e))
		panic(e)
	}
}

func (a *App) InitFiles() {
	a.ensureFile(a.LogFile)
}

func (a *App) ensureFile(s string) {
	f, err := os.Create(s)
	a.check(err, "creating "+s)
	f.Close()
	return
}

func (a *App) WriteLog(s string) {
	f, err := os.OpenFile(a.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	defer f.Close()
	if err != nil {
		panic(fmt.Sprintf("error opening %s\n", a.LogFile))
	}
	_, err = f.WriteString(s)
	if err != nil {
		panic(fmt.Sprintf("error writing to %s\n", a.LogFile))
	}
	return
}

func (a *App) Write(p []byte) (int, error) {
	a.CurrentText += string(p)
	a.SshText.SetText(a.CurrentText, true)
	a.SetFocus(a.SshText)
	return len(p), nil
}

func (a *App) Tick() {
	ticker := time.NewTicker(50 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				if a.SshOpen {
					// f, err := os.Open(a.LogFile)
					// if err != nil {
					// 	panic(err)
					// }
					// data, err := io.ReadAll(f)
					// if err != nil {
					// 	panic(err)
					// }
					// a.SshText.SetText(string(data), true)

				} else {
					// a.WriteLog("empty outpipe\n")
				}
			}
		}
	}()

}

// func (a *App) Tick() {
// 	ticker := time.NewTicker(20 * time.Millisecond)
// 	quit := make(chan struct{})
// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				if a.OutPipe != nil {
// 					fmt.Println("reading outpipe")
// 					a.WriteLog("checking outpipe...")
// 					res, err := io.ReadAll(a.OutPipe)
// 					a.check(err, "reading from outpipe")
// 					a.WriteLog("response: %s" + string(res))
// 				} else {
// 					a.WriteLog("empty outpipe")
// 				}
// 				return
// 			case <-quit:
// 				ticker.Stop()
// 				return
// 			}
// 		}
// 	}()
// }
//
//
// func (a *App) Write(p []byte) (int, error) {
// 	return len(p), nil
// }

func (a *App) Read(p []byte) (n int, err error) {
	a.CurrentText += string(p)
	a.SshText.SetText(a.CurrentText, true)
	a.SetFocus(a.SshText)
	return len(p), nil
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
// var text strings.Builder

// type MyWriter struct{}
//
// // Write(p []byte) (n int, err error)
//
// func (w *MyWriter) Write(p []byte) (int, error) {
// 	fmt.Printf("p: %s\n", string(p))
// 	return len(p), nil
// }

// var MyApp *App
//
// MyApp := &tui.App{
// 	Application: tview.NewApplication(),
// 	Layout:      tui.NewMyGrid(),
// 	InnerGrid:   tui.NewMyGrid(),
// 	InnerLeft:   tui.NewMyGrid(),
// 	SshWrites:   "",
// }

// if event.Key() == tcell.KeyEnter {
// 	fmt.Printf("cmdStack: %s\n", string(MyApp.cmdStack))
// 	MyApp.cmdStack = []byte{}
// 	cmds := strings.Split(string(MyApp.cmdStack), " ")
// 	// MyApp.cmd = exec.Command(cmds[0], cmds[1:]...)
// 	MyApp.Suspend(func() {
// 		go func() {
// 			cmd := exec.Command(cmds[0], cmds[1:]...)
// 			cmd.Stdout = MyApp
// 			cmd.Stderr = MyApp
// 			cmd.Stdin = MyApp
// 			err := cmd.Run()
// 			if err != nil {
// 				fmt.Printf("error: %s \n", err)
// 			}
// 		}()
// 	})
// }

// func (a *App) ReadSshIn() string {
//
//		// f, err := os.Open(a.SshIn.name)
//		// defer f.Close()
//		// if err != nil {
//		// 	panic(fmt.Sprintf("error opening %s: %s\n", a.SshIn, err))
//		// }
//		//
//		// text, err := io.ReadAll(f)
//		// if err != nil {
//		// 	panic(fmt.Sprintf("error reading all from %s: %s\n", a.SshIn, err))
//		// }
//		//
//		// return string(text)
//	}
//
// func (a *App) ReadOutpipe() {
// 	if a.OutPipe == nil {
// 		return
// 	}
// 	defer a.OutPipe.Close()
// 	data, err := io.ReadAll(a.OutPipe)
// 	// var data []byte
// 	// _, err := a.OutPipe.Read(data)
// 	// a.WriteLog(string(data))
// 	a.check(fmt.Errorf("reading from outpipe: %s\n", err), "")
//
// 	a.CurrentText += string(data)
// 	a.SshText.SetText(a.CurrentText, true)
// 	a.SetFocus(a.SshText)
//
// }
/*
		if event.Key() == tcell.KeyCtrlA && !sshopen {
			sshopen = true
			MyApp.Suspend(func() {
				go func() {
					cmd := MyApp.cmd
					var err error
					cmd.Stdin = MyApp.SshIn
					cmd.Stdout = MyApp.SshOut

					err = cmd.Run()
					MyApp.check(err, "running cmd ")
				}()
				return
			})
			return event
		}

		if sshopen {
			MyApp.SshIn.Write([]byte{byte(event.Rune())})
			// MyApp.SshOut.Write([]byte{byte(event.Rune())})
		}
		return event
	})
*/
// func (a *App) Write(p []byte) (int, error) {
// 	a.CurrentText += string(p)
// 	a.WriteLog(string(p))
// 	a.SshText.SetText(a.CurrentText, true)
// 	a.SetFocus(a.SshText)
// 	return 0, nil
// }
// a.InPipe = stdin
// a.OutPipe = stdout
// if event.Key() == tcell.KeyEnter {
// 	// MyApp.RunCmd()
// 	// MyApp.session.Wait()
// 	data, err := io.ReadAll(MyApp.OutPipe)
// 	if err != nil {
// 		panic(err)
// 	}
// 	MyApp.WriteLog(string(data))
// } else {
// 	MyApp.CurrentText += string(event.Rune())
// }
// io.WriterTo

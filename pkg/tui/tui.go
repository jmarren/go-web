package tui

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.org/x/crypto/ssh"
)

type App struct {
	*tview.Application
	Layout      *MyGrid
	InnerGrid   *MyGrid
	InnerLeft   *MyGrid
	SshWrites   *strings.Builder
	SshText     *tview.TextArea
	CurrentText string
	SshOpen     bool
	InPipe      io.WriteCloser
	LogFile     string
	OutPipe     io.Reader
	cmd         *exec.Cmd
	SshIn       *File
	SshOut      *outFile
	cmdStack    []byte
	session     *ssh.Session
	sshClient   *ssh.Client
}

type outFile struct {
	name     string
	textArea *tview.TextArea
}

func (f *outFile) Write(p []byte) (int, error) {
	file, err := os.OpenFile(f.name, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	defer file.Close()
	if err != nil {
		panic(fmt.Sprintf("error opening %s\n", f.name))
	}
	_, err = file.WriteString(string(p))
	if err != nil {
		panic(fmt.Sprintf("error writing to %s\n", f.name))
	}

	return len(p), nil
}

func newOutFile(name string) *outFile {
	return &outFile{
		name: name,
	}
}

type File struct {
	name string
}

// func (f *File) Read(p []byte) (int, error) {
// 	file, err := os.Open(f.name)
// 	defer file.Close()
// 	if err != nil {
// 		panic(fmt.Sprintf("error opening %s: %s\n", f.name, err))
// 	}
// 	n, err := file.Read(p)
// 	if err != nil && err != io.EOF {
// 		panic(fmt.Sprintf("error reading from %s: %s\n", f.name, err))
// 	}
// 	return n, nil
//
// }

func (f *File) Write(p []byte) (int, error) {
	file, err := os.OpenFile(f.name, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	defer file.Close()
	if err != nil {
		panic(fmt.Sprintf("error opening %s\n", f.name))
	}
	_, err = file.WriteString(string(p))
	if err != nil {
		panic(fmt.Sprintf("error writing to %s\n", f.name))
	}

	return len(p), nil
}

func newFile(s string) *File {
	return &File{
		name: s,
	}
}

func Create() *App {

	// func OpenFile(name string, flag int, perm FileMode) (*File, error) {
	// sshIn.
	MyApp := &App{
		Application: tview.NewApplication(),
		Layout:      NewMyGrid(),
		InnerGrid:   NewMyGrid(),
		InnerLeft:   NewMyGrid(),
		LogFile:     "app.log",
		SshWrites:   new(strings.Builder),
		SshOpen:     false,
		SshIn:       newFile("sshIn.txt"),
		SshOut:      newOutFile("sshOut.txt"),
		cmd:         exec.Command("myapp", "connect", "devdb"),
		CurrentText: "",
	}

	MyApp.InitFiles()

	MyApp.WriteLog("hi")

	// MyApp.cmd.Stdout = MyApp.SshWrites

	MyApp.CreateSshTextArea()

	MyApp.SshOut.textArea = MyApp.SshText

	MyApp.CreateInnerLeft()

	MyApp.CreateInnerGrid()

	MyApp.CreateLayout()

	MyApp.SetRoot(MyApp.Layout, true).SetFocus(MyApp.Layout)

	var sshopen = false

	MyApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlA && !sshopen {
			MyApp.SetClient()
			sshopen = true
			return event
		}
		if sshopen {
			if event.Key() == tcell.KeyEnter {
				MyApp.RunCmd()
			} else {
				MyApp.CurrentText += string(event.Rune())
			}
		}
		return event
	})
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

	return MyApp
}

func (a *App) Write(p []byte) (int, error) {
	a.CurrentText += string(p)
	a.WriteLog(string(p))
	// a.SshText.Replace(0, len(a.CurrentText), newText)
	// a.CurrentText = newText
	a.SshText.SetText(a.CurrentText, true)
	a.SetFocus(a.SshText)
	return 0, nil
}

func (a *App) check(e error, s string) {
	if e != nil {
		a.WriteLog(fmt.Sprintf("error %s:  %s", s, e))
		panic(e)
	}
}

func (a *App) InitFiles() {
	a.ensureFile(a.SshIn.name)
	a.ensureFile(a.SshOut.name)
	a.ensureFile(a.LogFile)
}

func (a *App) ensureFile(s string) {
	f, err := os.Create(s)
	a.check(err, "creating "+s)
	f.Close()
	return
}

func (a *App) WriteSshIn(p []byte) {
	sshIn, err := os.OpenFile(a.SshIn.name, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModeAppend)
	defer sshIn.Close()
	a.check(err, "opening "+a.SshIn.name)
	_, err = sshIn.Write(p)
	a.check(err, "writing to "+a.SshIn.name)
	a.CurrentText += string(p)
	a.SshText.SetText(a.CurrentText, true)
	a.SetFocus(a.SshText)

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

//
// func (a *App) Write(p []byte) (int, error) {
// 	return len(p), nil
// }

func (a *App) Read(p []byte) (n int, err error) {
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
// 	// f, err := os.Open(a.SshIn.name)
// 	// defer f.Close()
// 	// if err != nil {
// 	// 	panic(fmt.Sprintf("error opening %s: %s\n", a.SshIn, err))
// 	// }
// 	//
// 	// text, err := io.ReadAll(f)
// 	// if err != nil {
// 	// 	panic(fmt.Sprintf("error reading all from %s: %s\n", a.SshIn, err))
// 	// }
// 	//
// 	// return string(text)
// }
// func (a *App) Tick() {
// 	ticker := time.NewTicker(100 * time.Millisecond)
// 	quit := make(chan struct{})
// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				a.ReadOutpipe()
// 				return
// 			case <-quit:
// 				ticker.Stop()
// 				return
// 			}
// 		}
// 	}()
// }
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

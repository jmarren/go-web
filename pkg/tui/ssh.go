package tui

import (
	"log"

	"golang.org/x/crypto/ssh"
)

func (a *App) SetClient() {
	config := &ssh.ClientConfig{
		User: "test",
		Auth: []ssh.AuthMethod{
			ssh.Password("test"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "127.0.0.1:200", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	a.sshClient = client
	a.SetSession()
}

func (a *App) SetSession() {

	newSession, err := a.sshClient.NewSession()
	a.check(err, "creating ssh session")
	a.session = newSession
	// err = a.session.Shell()
	// if err != nil {
	// 	a.WriteLog("err running Shell(): " + err.Error())
	// }

	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// Request pseudo terminal
	if err := a.session.RequestPty("bash", 40, 80, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}
	a.InPipe, _ = a.session.StdinPipe()
	// a.session.Stdin = a
	a.session.Stdout = a

	// Start remote shell
	if err := a.session.Shell(); err != nil {
		log.Fatal("failed to start shell: ", err)
	}

}

// func (a *App) RunCmd() {
// 	ok, err := a.session.SendRequest("bash", true, []byte(a.CurrentText))
// 	if err != nil {
// 		panic(err)
// 	}
// 	if !ok {
// 		fmt.Println("pty-req not ok")
// 	}
// 	// a.SetSession()
// 	a.CurrentText = ""
//
// }

// func (a *App) RunCmd() {
//
// 	defer a.ResetCmd()
//
// 	a.WriteLog("running command: " + a.CurrentText + "\n")
//
// 	if a.CurrentText == "" {
// 		a.SshText.SetText("speak up", true)
// 		return
// 	}
//
// 	// Run the command on the remote machine
// 	if ok, err := a.session.SendRequest("bash", true, []byte(a.CurrentText)); err != nil {
// 		a.SshText.SetText("huh? "+err.Error(), true)
// 		if !ok {
// 			a.WriteLog("send request not ok")
// 		}
// 		return
// 	}
//
// 	if err := a.session.Wait(); err != nil {
// 		a.SshText.SetText("huh?"+err.Error(), true)
// 		return
// 	}
//
// 	res, err := io.ReadAll(a.OutPipe)
//
// 	a.check(err, "reading outpipe")
//
// 	a.SshText.SetText(string(res), true)
//
// 	a.WriteLog("response: " + string(res) + "\n")
//
// }

// func (a *App) Connect() {
// 	config := &ssh.ClientConfig{
// 		User: "test",
// 		Auth: []ssh.AuthMethod{
// 			ssh.Password("test"),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 	}
// 	client, err := ssh.Dial("tcp", "127.0.0.1:200", config)
// 	if err != nil {
// 		log.Fatal("Failed to dial: ", err)
// 	}
// 	// defer client.Close()
//
// 	// Each ClientConn can support multiple interactive sessions,
// 	// represented by a Session.
// 	session, err := client.NewSession()
// 	if err != nil {
// 		log.Fatal("Failed to create session: ", err)
// 	}
//
// 	// session.Stdout = a.SshOut
// 	// session.Stdin = os.Stdin
// 	//
// 	// session.Start("echo hi")
//
// 	a.session = session
//
// 	stdout, err := a.session.StdoutPipe()
// 	if err != nil {
// 		log.Fatalf("failed to get stdout: %s", err)
// 	}
//
// 	a.OutPipe = stdout
// 	// defer session.Close()
//
// }

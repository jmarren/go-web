package tui

import (
	"io"
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
	stdout, err := a.session.StdoutPipe()
	if err != nil {
		log.Fatalf("failed to get stdout: %s", err)
	}

	a.OutPipe = stdout
}

func (a *App) ResetCmd() {
	a.SetSession()
	a.CurrentText = ""

}

func (a *App) RunCmd() {

	defer a.ResetCmd()

	a.WriteLog("running command: " + a.CurrentText + "\n")

	if a.CurrentText == "" {
		a.SshText.SetText("speak up", true)
		return
	}

	// Run the command on the remote machine
	if err := a.session.Start(a.CurrentText); err != nil {
		a.SshText.SetText("huh? "+err.Error(), true)
		return
	}
	if err := a.session.Wait(); err != nil {
		a.SshText.SetText("huh?"+err.Error(), true)
		return
	}

	res, err := io.ReadAll(a.OutPipe)

	a.check(err, "reading outpipe")

	a.SshText.SetText(string(res), true)

	a.WriteLog("response: " + string(res) + "\n")

}

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

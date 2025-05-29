package myssh

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func Connect() {
	// func NewClientConn(c net.Conn, addr string, config *ClientConfig) (Conn, <-chan NewChannel, <-chan *Request, error) {
	// conn := ssh.NewClientConn()
	// conn, err := net.Dial("tcp", "golang.org:80")
	// if err != nil {
	// 	// handle error
	// }
	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// status, err := bufio.NewReader(conn).ReadString('\n')
	// ...

	// var hostKey ssh.PublicKey
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
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
	defer client.Close()

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

}

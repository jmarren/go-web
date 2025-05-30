package myssh

import "golang.org/x/crypto/ssh"

type IInstance interface {
	ClientConfig() *ssh.ClientConfig
	Network() string
	Addr() string
}

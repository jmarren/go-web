package myssh

import (
	"errors"
	// "io"

	"golang.org/x/crypto/ssh"
)

type TermConfig struct {
	clientCfg *ssh.ClientConfig
	network   string
	addr      string
	// writer    io.Writer
}

func (t *TermConfig) error(e ...error) error {
	e = append(e, ErrTermConfigValidate)
	return errors.Join(e...)
}

func (t *TermConfig) Validate() error {
	if t.addr == "" {
		return t.error(ErrAddrNotSet)
	}
	if t.network == "" {
		return t.error(ErrNetworkNotSet)
	}
	if t.clientCfg == nil {
		return t.error(ErrClientConfigValiate)
	}
	// if t.writer == nil {
	// 	return t.error(ErrWriterNotSet)
	// }
	return nil
}

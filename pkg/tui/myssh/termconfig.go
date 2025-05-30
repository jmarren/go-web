package myssh

import (
	"errors"
	// "io"

	"golang.org/x/crypto/ssh"
)

type TermConfig struct {
	ClientCfg *ssh.ClientConfig
	Network   string
	Addr      string
}

func (t *TermConfig) error(e ...error) error {
	e = append(e, ErrTermConfigValidate)
	return errors.Join(e...)
}

// func (t *TermConfig) Addr() string {
// 	return t.addr
// }
//
// func (t *TermConfig) Network() string {
// 	return t.network
// }
//
// func (t *TermConfig) ClientConfig() *ssh.ClientConfig {
// 	return t.clientCfg
// }

// func (t *TermConfig) Validate() error {
// 	if t.addr == "" {
// 		return t.error(ErrAddrNotSet)
// 	}
// 	if t.network == "" {
// 		return t.error(ErrNetworkNotSet)
// 	}
// 	if t.clientCfg == nil {
// 		return t.error(ErrClientConfigValiate)
// 	}
// 	// if t.writer == nil {
// 	// 	return t.error(ErrWriterNotSet)
// 	// }
// 	return nil
// }

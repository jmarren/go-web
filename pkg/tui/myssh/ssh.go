package myssh

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"io"
)

var modes = ssh.TerminalModes{
	ssh.ECHO:          0,     // disable echoing
	ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
	ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
}

type SshService struct {
	client     *ssh.Client
	session    *ssh.Session
	inPipe     io.WriteCloser
	termConfig *TermConfig
	Active     bool
	writer     io.Writer
}

func New(writer io.Writer) *SshService {
	return &SshService{
		writer: writer,
	}
}

func (s *SshService) error(e ...error) error {
	e = append(e, ErrSsh)
	return errors.Join(e...)
}

// sets the client using the provided termConfig
func (s *SshService) setClient() error {
	// err if termConfig not set
	if s.termConfig == nil {
		return s.error(ErrTermConfigNotSet)
	}

	t := s.termConfig

	// dial using termConfig
	client, err := ssh.Dial(t.network, t.addr, t.clientCfg)
	if err != nil {
		return s.error(ErrDialFailed, err)
	}

	// set client
	s.client = client

	// call setSession
	if err = s.setSession(); err != nil {
		return s.error(ErrSetSession, err)
	}

	return nil
}

func (s *SshService) setSession() error {
	// create session
	newSession, err := s.client.NewSession()
	if err != nil {
		return s.error(err, ErrCreateSession)
	}
	// set session
	s.session = newSession

	// Request pseudo terminal
	if err := s.session.RequestPty("bash", 40, 80, modes); err != nil {
		return s.error(err, ErrReqPTY)
	}

	// set inPipe to session StdinPip
	s.inPipe, _ = s.session.StdinPipe()

	// set session Stout to writer supplied by termConfig
	s.session.Stdout = s.writer

	// Start remote shell
	if err := s.session.Shell(); err != nil {
		return s.error(err, ErrShellStart)
	}
	return nil
}

// This method can be used by the caller to pass data to the remote terminal
func (s *SshService) PipeIn(p []byte) error {
	if s.inPipe == nil {
		return ErrInpipeNotSet
	}
	s.inPipe.Write(p)
	return nil
}

// this is where the service will use the provided instance name
// to construct a termConfig which it will set as s.termConfig
func (s *SshService) Connect(instance string) error {
	switch instance {
	case "devdb":
		s.termConfig = &TermConfig{
			clientCfg: &ssh.ClientConfig{
				User: "test",
				Auth: []ssh.AuthMethod{
					ssh.Password("test"),
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			},
			network: "tcp",
			addr:    "127.0.0.1:200",
		}
	}

	if err := s.termConfig.Validate(); err != nil {
		return err
	}

	if err := s.setClient(); err != nil {
		return s.error(ErrConn, err)
	}
	s.Active = true
	return nil
}

// config := &ssh.ClientConfig{
// 	User: "test",
// 	Auth: []ssh.AuthMethod{
// 		ssh.Password("test"),
// 	},
// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// }

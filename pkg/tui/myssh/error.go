package myssh

import "errors"

var ErrDialFailed = errors.New("Failed to Dial")
var ErrCreateSession = errors.New("Error creating session")
var ErrReqPTY = errors.New("Error requesting pty terminal")
var ErrShellStart = errors.New("Error starting shell")
var ErrInpipeNotSet = errors.New("Inpipe is not set")
var ErrSetSession = errors.New("Error setting ssh.session")
var ErrSsh = errors.New("sshService Error")
var ErrTermConfigNotSet = errors.New("sshService.termConfig not set")
var ErrConn = errors.New("connection error")

// term config
var ErrTermConfigValidate = errors.New("TermConfig error")
var ErrClientConfigValiate = errors.New("client config error")
var ErrNetworkNotSet = errors.New("Network not set")
var ErrAddrNotSet = errors.New("Addr not set")
var ErrWriterNotSet = errors.New("Writer notset")

// network   string
// addr      string
// writer    io.Writer

package instance

import "io"

type sshService interface {
	Connect(instance string, writer io.Writer) error
}

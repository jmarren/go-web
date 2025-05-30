package instance

import (
	"fmt"
	"io"
	"time"

	"github.com/jmarren/go-web/pkg/tui/myssh"
)

type Instance struct {
	Name      string
	Online    bool
	StartTime time.Time
	*myssh.TermConfig
}

func (i *Instance) uptimeStr() string {
	uptime := time.Since(i.StartTime)
	hours := int(uptime.Round(time.Hour).Hours())
	mins := int(uptime.Round(time.Minute).Minutes()) - 60*hours
	return fmt.Sprintf("%dh %dm", hours, mins)
}

func (i *Instance) Connect(s *myssh.SshService, w io.Writer) {
	s.Connect(i.TermConfig, w)
}

func (i *Instance) status() string {
	if i.Online {
		return "online"
	}
	return "offline"
}

func (i *Instance) TableRow() []string {
	return []string{
		i.Name,
		i.status(),
		i.uptimeStr(),
		"db",
	}
}

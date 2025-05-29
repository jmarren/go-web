package logger

import (
	"errors"
	"os"
)

type Logger struct {
	filepath string
}

func New(filepath string) (*Logger, error) {
	l := &Logger{
		filepath,
	}
	if err := l.ensureFile(); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *Logger) error(e ...error) error {
	e = append(e, ErrLogger)
	return errors.Join(e...)
}

func (l *Logger) ensureFile() error {
	f, err := os.Create(l.filepath)
	defer f.Close()
	if err != nil {
		return l.error(ErrFileCreate)
	}
	return nil
}

func (l *Logger) Write(s string) error {
	f, err := os.OpenFile(l.filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	defer f.Close()
	if err != nil {
		return l.error(ErrFileOpen)
	}
	_, err = f.WriteString(s)
	if err != nil {
		return l.error(ErrFileWrite)
	}
	return nil
}

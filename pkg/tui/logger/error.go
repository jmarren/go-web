package logger

import "errors"

var ErrLogger = errors.New("Logger Error")
var ErrFile = errors.New("file error")
var ErrFileCreate = errors.Join(ErrFile, errors.New("ensuring/creating file"))
var ErrFileOpen = errors.Join(ErrFile, errors.New("opening file"))
var ErrFileWrite = errors.Join(ErrFile, errors.New("writing to file"))

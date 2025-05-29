package tui

import "errors"

var ErrFile = errors.New("file error")
var ErrFileCreate = errors.Join(ErrFile, errors.New("ensuring/creating file"))
var ErrFileOpen = errors.Join(ErrFile, errors.New("opening file"))

package utils

import "io"

func PanicE(e error) {
	if e != nil {
		panic(e)
	}
}

type Writer struct {
	writeFunc func(p []byte)
}

func (w *Writer) Write(p []byte) (int, error) {
	w.writeFunc(p)
	return len(p), nil
}

func WriterFrom(writeFunc func(p []byte)) io.Writer {
	return &Writer{
		writeFunc: writeFunc,
	}
}

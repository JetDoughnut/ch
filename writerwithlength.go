package proxychannel

import (
	"io"
	"net/http"
)

// WriterWithLength .
type WriterWithLength struct {
	writer        interface{} // io.Writer or http.ResponseWriter
	interfaceType int
	length        int
}

func (w *WriterWithLength) Write(b []byte) (n int, err error) {
	if w.interfaceType == 0 {
		// http.ResponseWriter
		respWriter, ok := w.writer.(http.ResponseWriter)
		if !ok {
			panic("w.writer is not a http.ResponseWriter")
		}
		n, err = respWriter.Write(b)
		w.length += n
	} else {
		// io.Writer
		ioWriter, ok := w.writer.(io.Writer)
		if !ok {
			panic("w.writer is not a io.Writer")
		}
		n, err = ioWriter.Write(b)
		w.length += n
	}
	return
}

// Length .
func (w *WriterWithLength) Length() int {
	return w.length
}

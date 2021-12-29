package chain_of_responsibilities

import "io"

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (w *WriterLogger) Next(s string) {
	if w.Writer != nil {
		w.Writer.Write([]byte("WriterLogger: " + s))
	}
	if w.NextChain != nil {
		w.NextChain.Next(s)
	}

}

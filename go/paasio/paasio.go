package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	io.Reader
	data stats
}

type writeCounter struct {
	io.Writer
	data stats
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

type stats struct {
	ops   int
	bytes int
	mu    sync.Mutex
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{
		r,
		stats{},
	}
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.data.mu.Lock()
	defer rc.data.mu.Unlock()
	return int64(rc.data.bytes), rc.data.ops
}

func (rc *readCounter) Read(bytes []byte) (int, error) {
	rc.data.mu.Lock()
	defer rc.data.mu.Unlock()
	b, err := rc.Reader.Read(bytes)
	rc.data.bytes += b
	rc.data.ops++
	return b, err
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{
		w,
		stats{},
	}
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.data.mu.Lock()
	defer wc.data.mu.Unlock()
	return int64(wc.data.bytes), wc.data.ops
}

func (wc *writeCounter) Write(bytes []byte) (int, error) {
	wc.data.mu.Lock()
	defer wc.data.mu.Unlock()
	b, err := wc.Writer.Write(bytes)
	wc.data.bytes += b
	wc.data.ops++
	return b, err
}

func NewReadWriteCounter(rw io.ReadWriter) *readWriteCounter {
	return &readWriteCounter{
		NewReadCounter(rw),
		NewWriteCounter(rw),
	}
}

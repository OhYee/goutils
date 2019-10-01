package goutils

import (
	"fmt"
	"io"
	"sync"
)

var errwrite = fmt.Errorf("Pipe has closed")

// Pipe mutil thread supported write-read pipe buffer
type Pipe struct {
	data   chan byte
	close  chan bool
	closed bool
	mutex  *sync.Mutex
}

// NewPipe create a pipe
func NewPipe(length int, b ...[]byte) Pipe {
	p := Pipe{
		data:   make(chan byte, length),
		close:  make(chan bool, 1),
		closed: false,
		mutex:  &sync.Mutex{},
	}

	data := make([]byte, 0)
	for _, d := range b {
		data = append(data, d...)
	}
	p.Write(data)
	return p
}

// Read data from pipe
func (p Pipe) Read(b []byte) (n int, err error) {
	for i := 0; i < len(b); i++ {
		select {
		case d := <-p.data:
			b[i] = d
			n++
		default:
			if n == 0 && p.isClose() {
				err = io.EOF
			}
			break
		}
	}
	b = b[:n]
	return
}

// Write data to pipe
func (p Pipe) Write(b []byte) (n int, err error) {
	if p.isClose() {
		err = errwrite
		return
	}
	for _, bb := range b {
		p.data <- bb
		n++
	}
	return
}

func (p Pipe) isClose() bool {
	p.mutex.Lock()
	closed := p.closed
	p.mutex.Unlock()
	return closed
}

// Close the pipe
func (p *Pipe) Close() {
	p.mutex.Lock()
	p.closed = true
	p.mutex.Unlock()
}

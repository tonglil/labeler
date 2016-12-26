package logs

import (
	"fmt"
	"io"
	"runtime"
	"sync"
	"time"
)

const format = "2006/01/02 15:04:05"

type DetailedLogger struct {
	mu     sync.Mutex
	output io.Writer
}

func NewDetailedLogger(w io.Writer) *DetailedLogger {
	return &DetailedLogger{
		output: w,
	}
}

func (d *DetailedLogger) Write(p []byte) (n int, err error) {
	var file string
	var line int
	var ok bool

	_, file, line, ok = runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}

	now := time.Now().UTC().Format(format)
	out := fmt.Sprintf("%s %s:%d %s", now, file, line, p)

	d.mu.Lock()
	defer d.mu.Unlock()

	return d.output.Write([]byte(out))
}

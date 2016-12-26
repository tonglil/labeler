package logs

import (
	"fmt"
	"io"
)

type Log bool

var Output io.Writer
var Threshold int

func V(level int) Log {
	if Threshold >= level {
		return Log(true)
	}

	return Log(false)
}

func (l Log) Info(a ...interface{}) {
	if l {
		fmt.Fprintln(Output, a...)
	}
}

func (l Log) Infoln(a ...interface{}) {
	if l {
		fmt.Fprintln(Output, a...)
	}
}

func (l Log) Infof(f string, a ...interface{}) {
	if l {
		fmt.Fprintf(Output, f, a...)
		fmt.Fprintln(Output, "")
	}
}

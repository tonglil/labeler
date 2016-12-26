package main

import (
	"os"

	"github.com/tonglil/labeler/cmd"
	"github.com/tonglil/labeler/logs"
)

func init() {
	logs.Output = os.Stdout
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

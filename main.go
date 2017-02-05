package main

import (
	"os"

	"github.com/tonglil/labeler/cmd"
	"github.com/tonglil/labeler/logs"
	"github.com/tonglil/versioning"
)

var version string

func init() {
	versioning.Set(version)

	logs.Output = os.Stdout
}

func main() {
	cmd.Execute()
}

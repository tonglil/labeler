package main

import (
	"flag"
	"os"

	"github.com/golang/glog"
	"github.com/tonglil/labeler/cmd"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Set("v", "9")
}

func main() {
	flag.Parse()

	if err := cmd.Execute(); err != nil {
		glog.V(0).Info(err)
		os.Exit(1)
	}

	os.Exit(0)
}

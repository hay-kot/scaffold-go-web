package main

import (
	"fmt"

	"{{ .Scaffold.gomod }}/cmd/api/web"
)

var (
	version   = "nightly"
	commit    = "HEAD"
	buildTime = "now"
)

func build() string {
	short := commit
	if len(short) > 7 {
		short = short[:7]
	}

	return fmt.Sprintf("%s, commit %s, built at %s", version, short, buildTime)
}

func main() {
	conf, err := web.ConfigFromCLI()
	if err != nil {
		panic(err)
	}

	args := &web.WebArgs{
		Conf:  conf,
		Build: build(),
	}

	if err := web.New(args).Start(); err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/mdelapenya/cauldrongo/cmd"
)

// nolint: gochecknoglobals
var (
	version = ""
	commit  = ""
	date    = ""
	builtBy = ""
)

func main() {
	cmd.Execute()
}

package main

import (
	"github.com/charmbracelet/log"
	"github.com/zcubbs/grill/cmd/cli/cmd"
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

func init() {
	cmd.Version = Version
	cmd.Commit = Commit
	cmd.Date = Date
}

func main() {
	log.SetReportTimestamp(false)
	cmd.Execute()
}

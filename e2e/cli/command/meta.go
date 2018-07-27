package command

import (
	"flag"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/mitchellh/cli"
)

type Meta struct {
	Ui     cli.Ui
	logger hclog.Logger

	verbose bool
}

func (m *Meta) FlagSet(n string) *flag.FlagSet {
	f := flag.NewFlagSet(n, flag.ContinueOnError)

	f.BoolVar(&m.verbose, "v", false, "Toggle verbose output")
	return f
}

package config

import (
	"github.com/spf13/cobra"
	"kiririx/gox/cmd/config/act"
)

func init() {
	CmdConfig.AddCommand(act.CmdLs, act.CmdRm, act.CmdSet)
}

var CmdConfig = &cobra.Command{
	Use: "config",
}

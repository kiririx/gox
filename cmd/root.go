/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	addPkg "kiririx/gox/cmd/add"
	"kiririx/gox/cmd/config"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "gox",
}

var lib = ""
var act = ""

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetUsageTemplate(help)
	_ = rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(config.CmdConfig)
	rootCmd.AddCommand(addPkg.CmdAdd)
}

package main

import (
	"github.com/overbool/packr"
	"github.com/spf13/cobra"
)

var all bool

func init() {
	versionCMD.Flags().BoolVar(&all, "all", false, "show all version info")
}

var versionCMD = &cobra.Command{
	Use:   "version [flags]",
	Short: "Show version about app",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Packr version: %s-%s\n", packr.CurrentVersion, packr.CurrentCommit)
		if all {
			cmd.Printf("App build date: %s\n", packr.BuildDate)
			cmd.Printf("System version: %s\n", packr.Platform)
			cmd.Printf("Golang version: %s\n", packr.GoVersion)
		}
	},
}

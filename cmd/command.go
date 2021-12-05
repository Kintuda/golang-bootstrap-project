package cmd

import "github.com/spf13/cobra"

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use: "app",
}

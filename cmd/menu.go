package cmd

import (
	"github.com/ohbyeongmin/vscodeCLI/cli"
	"github.com/spf13/cobra"
)


var menuCmd = &cobra.Command{
	Use: "menu",
	Short: "Show menu",
	Run: func(cmd *cobra.Command, args []string) {
		cli.Start()
	},
}

func init() {
	rootCmd.AddCommand(menuCmd)
}
package cmd

import (
	"fmt"

	"github.com/ohbyeongmin/vscodeCLI/cli"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vscodeCLI",
	Short: "Run project",
	Long: `vscodeCLI makes it easy to find and run projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ :=cmd.Flags().GetString("direct")
		err := cli.DirectExecProject(path)
		if err != nil {
			fmt.Printf("\nNot Exist %s, Please check Name or Path (If you are a first-time user, please enter the menu)\n\n\n",path)
			cmd.Help()
		}
	},
}

func Execute() {
	cli.InitializeApp()
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP("direct", "d", "", "Open Project directly, You should be know project name exactly")
}


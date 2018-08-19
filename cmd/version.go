package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + AppName,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

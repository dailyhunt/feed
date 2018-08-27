package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func versionCommand(appName string, version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of " + appName,
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
}

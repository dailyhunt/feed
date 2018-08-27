package cmd

import (
	"github.com/spf13/cobra"
	"github.com/dailyhunt/feed/server"
	"github.com/spf13/viper"
)

func serverCommand(appName string, version string, builder server.EndpointsBuilder) (*cobra.Command) {
	var port = viper.GetInt("server.port")

	return &cobra.Command{
		Use:   "server",
		Short: "Runs the server",
		Long:  ``,
		Run:   server.HttpServerRunner(appName, version, builder, port),
	}
}

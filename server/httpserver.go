package server

import (
	"fmt"
	"github.com/dailyhunt/feed/engine"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

type CobraRunner func(cmd *cobra.Command, args []string)

type EndpointsBuilder func() (endpoints []HttpEndpoint)

func HttpServerRunner(appName string, version string, builder EndpointsBuilder, port int) CobraRunner {
	return func(cmd *cobra.Command, args []string) {
		var router = gin.New()

		router.Use(ginrus.Ginrus(logger.StandardLogger(), time.RFC3339Nano, false))
		router.Use(gin.Recovery())

		var endpoints = builder()

		for _, endpoint := range endpoints {
			var e = engine.Build(endpoint.EngineDefn)

			logger.Info(fmt.Sprintf("Registering endpoint: METHOD=%s, PATH=%s", endpoint.Method, endpoint.Path))
			router.Handle(string(endpoint.Method), endpoint.Path, e.Serve)
		}

		// Do Stuff Here
		logger.Info(fmt.Sprintf("Starting %s server [%s] on port=%d", appName, version, port))
		router.Run(fmt.Sprintf(":%d", port))
	}
}

package cmd

import (
	"fmt"

	"github.com/dailyhunt/feed/server"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

func Execute(appName string, version string, builder server.EndpointsBuilder) {
	var rootCmd = &cobra.Command{
		Use:   "feed",
		Short: "feed",
		Long:  ``,
	}

	initEnv(rootCmd, appName)

	rootCmd.AddCommand(versionCommand(appName, version))
	rootCmd.AddCommand(serverCommand(appName, version, builder))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initEnv(rootCmd *cobra.Command, appName string) {
	var cfgFile string
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")

	//
	// enable commandline flags
	//
	viper.BindPFlags(pflag.CommandLine)

	//
	// enable environment
	//
	viper.SetEnvPrefix(appName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	initConfig(cfgFile, appName)
}

func defaultConfig() {
	//
	// viper defaults
	//
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("log.Level", "info")
}

func initConfig(cfgFile string, appName string) {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		//
		// Config file
		//
		env := os.Getenv("env")
		if env == "" {
			env = "local"
		}

		viper.SetConfigName(fmt.Sprintf("config-%s", env))
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("./" + appName + "/config")
		viper.AddConfigPath("$HOME/" + appName + "/config")
		viper.AddConfigPath("/usr/local/etc/" + appName + "/config")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error in reading config file: %s \n", err))
	}

	defaultConfig()

	configureLogger()
}

func configureLogger() {

	var logConfig logConfig

	err := viper.UnmarshalKey("log", &logConfig)
	if err != nil {
		panic(fmt.Errorf("Fatal error in reading 'log' config: %s \n", err))
	}

	if logConfig.Format == JSON {
		logger.SetFormatter(&logger.JSONFormatter{})
	} else if logConfig.Format == TEXT {
		logger.SetFormatter(&logger.TextFormatter{})
	} else {
		panic(fmt.Errorf("Unsupported log format: %s ! Supported Values: json, text \n", logConfig.Format))
	}

	if logConfig.Output == STDOUT {

		logger.SetOutput(os.Stdout)

	} else if logConfig.Output == FILE {
		logger.SetOutput(&lumberjack.Logger{
			Filename:   logConfig.File.Filename,
			MaxSize:    logConfig.File.MaxSize, // megabytes
			MaxBackups: logConfig.File.MaxBackups,
			MaxAge:     logConfig.File.MaxAge,   //days
			Compress:   logConfig.File.Compress, // disabled by default
		})
	} else {
		panic(fmt.Errorf("Unsupported log output: %s ! Supported Values: stdout, file \n", logConfig.Output))
	}

	logLevel, err := logger.ParseLevel(logConfig.Level)
	if err != nil {
		panic(fmt.Errorf("Unsupported log level: %s ! Supported Values: panic, fatal, error, warn, warning, debug, info \n", logConfig.Level))
	}

	logger.SetLevel(logLevel)
}

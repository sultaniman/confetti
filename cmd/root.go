package cmd

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

const BaseURL = "https://getout.cloud"
const DefaultPort = 4000

var (
	port    int
	rootCmd = cobra.Command{
		Use:   "getout [command]",
		Short: "getout",
		Long:  "getout",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Usage()
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	cobra.OnInitialize(configure)
	rootCmd.PersistentFlags().IntVar(&port, "port", DefaultPort, "server port")

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(testCmd)
}

func configure() {
	viper.SetDefault("base_url", BaseURL)
	viper.SetEnvPrefix("GO")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	logLevel, err := zerolog.ParseLevel(viper.GetString("log_level"))
	if err == nil {
		zerolog.SetGlobalLevel(logLevel)
	}
}

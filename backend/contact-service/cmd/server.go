package cmd

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/webtech-fmi/phonebook/backend/contact-service/pkg/configuration"
	"github.com/webtech-fmi/phonebook/backend/contact-service/pkg/infrastructure/log"
	"github.com/webtech-fmi/phonebook/backend/contact-service/pkg/web"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Launch the HTTP server.",
	Run: func(cmd *cobra.Command, args []string) {
		appConfiguration := &configuration.AppConfiguration{}
		absoluteConfigPath, err := filepath.Abs(configPath)
		if err != nil {
			fmt.Printf("error at establishing configuration path: [%s]", err.Error())
			os.Exit(1)
		}

		err = configuration.LoadYAML(appConfiguration, &absoluteConfigPath, nil, []string{"port", "log"})
		if err != nil {
			fmt.Printf("error at configuration loading: [%s]", err.Error())
			os.Exit(1)
		}

		f := os.Stdout
		if appConfiguration.LogLevel == "error" {
			f = os.Stderr
		}

		err = appConfiguration.Validate()
		if err != nil {
			fmt.Printf("error at configuration validation: [%s]", err.Error())
			os.Exit(1)
		}

		logger, err := log.NewZerolog(f, appConfiguration.LogLevel)
		if err != nil {
			fmt.Printf("could not instantiate zerolog: [%s]", err.Error())
			os.Exit(1)
		}

		if err = web.LaunchServer(appConfiguration, logger); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("Error launching privacy service webserver")
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

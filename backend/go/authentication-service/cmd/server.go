package cmd

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/configuration"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/web"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
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

		err = appConfiguration.Validate()
		if err != nil {
			fmt.Printf("error at configuration validation: [%s]", err.Error())
			os.Exit(1)
		}

		logger := log.NewZerolog(true)
		if err != nil {
			fmt.Printf("could not instantiate zerolog: [%s]", err.Error())
			os.Exit(1)
		}

		if err = web.LaunchServer(appConfiguration, logger); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("Error launching authentication service webserver")
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

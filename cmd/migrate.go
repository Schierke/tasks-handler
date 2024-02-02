package cmd

import (
	"github.com/Schierke/tasks-handler/config"
	"github.com/Schierke/tasks-handler/pkg/migrate"
	"github.com/Schierke/tasks-handler/pkg/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// TODO: move it to config file
func init() {
	rootCmd.AddCommand(hydrateCmd)
}

var hydrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrating database when starting in a new env",
	RunE: func(_ *cobra.Command, args []string) error {

		configFile := utils.LoadConfigFile()
		cfg, err := config.LoadAppConfig(configFile)
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to load configuration")
		}

		migrate.Migrate(cfg)

		return nil
	},
}

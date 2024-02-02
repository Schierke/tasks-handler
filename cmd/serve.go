package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Schierke/tasks-handler/config"
	"github.com/Schierke/tasks-handler/internals/handler"
	"github.com/Schierke/tasks-handler/internals/repository"
	"github.com/Schierke/tasks-handler/internals/service"
	"github.com/Schierke/tasks-handler/pkg/mongodb"
	"github.com/Schierke/tasks-handler/pkg/utils"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server and listen to HTTP Request",
	RunE: func(command *cobra.Command, args []string) error {
		configFile := utils.LoadConfigFile()
		appConfig, _ := config.LoadAppConfig(configFile)

		// setting up DB
		mongoClient, err := mongodb.SetupDB(appConfig)
		if err != nil {
			log.Fatal().Msg("Can't setup connection with DB. Abort")
			return err
		}
		database := mongoClient.Database(appConfig.Mongo.Name)

		// setting up chi router
		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Use(middleware.Timeout(120 * time.Second))

		// setting up handlers
		taskRepo := repository.NewTaskRepo(database)
		userRepo := repository.NewTaskRepo(database)
		shiftRepo := repository.NewtShiftRepo(database)
		slotRepo := repository.NewtSlotRepo(database)

		taskService := service.NewTaskService(taskRepo)
		userService := service.NewTaskService(userRepo)
		shiftService := service.NewShiftService(shiftRepo, slotRepo)

		handler.NewUserHandler(userService, r)
		handler.NewTaskHandler(taskService, r)
		handler.NewShiftHandler(shiftService, r)

		http.ListenAndServe(fmt.Sprintf(":%s", appConfig.Server.AppPort), r)
		return nil
	},
}

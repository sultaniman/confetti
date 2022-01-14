package cmd

import (
	"fmt"
	"github.com/imanhodjaev/getout/platform/db"
	"github.com/imanhodjaev/getout/platform/handlers"
	"github.com/imanhodjaev/getout/platform/keys"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Long:  "Start server",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := db.Connect(viper.GetString("db_uri"))
		if err != nil {
			return err
		}

		keyLoader := keys.GetLoader(viper.GetString("key_loader"))
		key, err := keyLoader.Load(viper.GetString("private_key"))
		if err != nil {
			return err
		}

		handler, err := handlers.NewHandler(db, key)
		if err != nil {
			return err
		}

		app := handlers.App(handler)
		return app.Listen(fmt.Sprintf(":%d", port))
	},
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sultaniman/confetti/platform/db"
	"github.com/sultaniman/confetti/platform/handlers"
	"github.com/sultaniman/confetti/platform/keys"
	"time"
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

		maxOpen := viper.GetInt("database_max_open")
		maxIdle := viper.GetInt("database_max_idle")

		fmt.Printf("DB Pool options: max_open=%d, max_idle=%d, connection ttl=1h\n", maxOpen, maxIdle)
		db.SetMaxIdleConns(maxIdle)
		db.SetMaxOpenConns(maxOpen)
		db.SetConnMaxLifetime(time.Hour)

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

func init() {
	viper.SetDefault("db_uri", "postgres://postgres:postgres@localhost:5432/confetti?sslmode=disable")
	viper.SetDefault("database_max_open", 50)
	viper.SetDefault("database_max_idle", 20)
	viper.SetDefault("private_key", "")
	viper.SetDefault("refresh_token_ttl", "4320h") // 180 days
	viper.SetDefault("access_token_ttl", "1h")     // 1 hour
	viper.SetDefault("mailer", "dummy")
	viper.SetDefault("from_email", "no-reply@secura.team")
	viper.SetDefault("verbose", false)
}

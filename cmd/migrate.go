package cmd

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/url"
	"strings"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply database migrations",
	Long:  "Apply database migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		dbUri := viper.GetString("db_uri")

		db, err := sql.Open("postgres", dbUri)
		if err != nil {
			return err
		}

		uri, err := url.Parse(dbUri)
		if err != nil {
			return err
		}

		driver, err := postgres.WithInstance(db, &postgres.Config{})
		m, err := migrate.NewWithDatabaseInstance(
			viper.GetString("migrations"),
			strings.TrimPrefix(uri.Path, "/"),
			driver,
		)

		if err != nil {
			return err
		}

		if err = db.Ping(); err != nil {
			return err
		}

		if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return err
		}

		return nil
	},
}

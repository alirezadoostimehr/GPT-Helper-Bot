package cmd

import (
	"context"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/config"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"regexp"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		migrateDB(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrateDB(ctx context.Context) {
	log.SetReportCaller(true)
	err := config.Load()
	if err != nil {
		log.Error(err.Error())
		return
	}

	conn, err := postgres.NewConnectionPool(ctx, config.GlobalConfig.Postgres)
	if err != nil {
		log.Error(err)
		return
	}
	defer conn.Close()

	migrations, err := getAppliedMigrations(ctx, conn)
	if err != nil && !regexp.MustCompile("relation \"migration\" does not exist").MatchString(err.Error()) {
		log.Error(err)
		return
	}

	folder, err := os.Open("./migrations")
	if err != nil {
		log.Error(err)
		return
	}
	defer folder.Close()

	files, err := folder.Readdir(0)
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range files {
		isApplied := false
		for _, m := range migrations {
			if v.Name() == m {
				isApplied = true
				break
			}
		}
		if !isApplied {
			log.Println("Applying migration", v.Name())
			dat, err := os.ReadFile("./migrations/" + v.Name())
			if err != nil {
				log.Error(err)
				return
			}
			_, err = conn.Exec(ctx, string(dat))
			if err != nil {
				log.Error(err)
				return
			}

			err = addMigration(ctx, conn, v.Name())
			if err != nil {
				log.Error(err)
				return
			}
		}
	}
}

func getAppliedMigrations(ctx context.Context, conn *postgres.ConnectionPool) ([]string, error) {
	query := `SELECT version FROM migration`

	migrations := make([]string, 0)

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return migrations, err
	}

	for rows.Next() {
		var version string
		err = rows.Scan(&version)
		if err != nil {
			return migrations, err
		}
		migrations = append(migrations, version)
	}

	return migrations, nil
}

func addMigration(ctx context.Context, conn *postgres.ConnectionPool, version string) error {
	query := `INSERT INTO migration (version) VALUES ($1)`

	_, err := conn.Exec(ctx, query, version)
	if err != nil {
		return err
	}

	return nil
}

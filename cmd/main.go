package main

import (
	"federated/config"
	"federated/internal/seeder"
	"federated/pkg/db/migrator"
	"federated/pkg/db/postgres"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	config, err := config.New("config/config.yml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	rootCmd := &cobra.Command{
		Use:   "federated-users",
		Short: "Генератор данных federated-users",
	}

	initMigrateCmd(rootCmd, config)
	initMigrateRollbackCmd(rootCmd, config)
	initSeedMakeCmd(rootCmd)

	err = rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initSeedMakeCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "seed:make",
		Short: "Создание сидера данных",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("\n")
			fmt.Printf("Make database csv seeder...\n")

			file, err := seeder.GenerateRows(500000)
			if err == nil {
				fmt.Printf("Generated file %s\n", file)
			} else {
				fmt.Printf("Generated error: %v", err)
			}
			fmt.Printf("\n")

			return err
		},
	}

	rootCmd.AddCommand(cmd)
}

func initMigrateCmd(rootCmd *cobra.Command, cfg config.Config) {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Применить миграции базы данных",
		RunE: func(cmd *cobra.Command, args []string) error {
			db := postgres.NewConnectOrFail(cfg.Postgres)
			mg := migrator.NewOrFail(db)

			fmt.Printf("\n")
			fmt.Printf("Run migrate database...\n")

			err := mg.Up()
			if err == nil {
				fmt.Printf("Migration complete\n")
			} else {
				fmt.Printf("Migration error: %v", err)
			}
			fmt.Printf("\n")

			return err
		},
	}

	rootCmd.AddCommand(cmd)
}

func initMigrateRollbackCmd(rootCmd *cobra.Command, cfg config.Config) {
	cmd := &cobra.Command{
		Use:   "migrate:rollback",
		Short: "Откатить миграции базы данных",
		RunE: func(cmd *cobra.Command, args []string) error {
			db := postgres.NewConnectOrFail(cfg.Postgres)
			mg := migrator.NewOrFail(db)

			fmt.Printf("\n")
			fmt.Printf("Run migrate rollback...\n")

			err := mg.Down()
			if err == nil {
				fmt.Printf("Migration complete\n")
			} else {
				fmt.Printf("Migration error: %v", err)
			}
			fmt.Printf("\n")

			return err
		},
	}

	rootCmd.AddCommand(cmd)
}

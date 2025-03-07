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
	initMigrateForceCmd(rootCmd, config)
	initMigrateRollbackCmd(rootCmd, config)
	initMigrateStatusCmd(rootCmd, config)
	initSeedMakeCmd(rootCmd)

	err = rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
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
				fmt.Printf("Migration: %v\n", err)
			}
			fmt.Printf("\n")

			return nil
		},
	}

	rootCmd.AddCommand(cmd)
}

func initMigrateForceCmd(rootCmd *cobra.Command, cfg config.Config) {
	cmd := &cobra.Command{
		Use:   "migrate:force",
		Short: "Применить миграции базы данных в состоянии dirty",
		RunE: func(cmd *cobra.Command, args []string) error {
			db := postgres.NewConnectOrFail(cfg.Postgres)
			mg := migrator.NewOrFail(db)

			fmt.Printf("\n")
			fmt.Printf("Run migrate force...\n")

			err := mg.Force()
			if err == nil {
				fmt.Printf("Migration complete\n")
			} else {
				fmt.Printf("Migration: %v\n", err)
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
				fmt.Printf("Rollback migration complete\n")
			} else {
				fmt.Printf("Rollback: %v\n", err)
			}
			fmt.Printf("\n")

			return nil
		},
	}

	rootCmd.AddCommand(cmd)
}

func initMigrateStatusCmd(rootCmd *cobra.Command, cfg config.Config) {
	cmd := &cobra.Command{
		Use:   "migrate:status",
		Short: "Состояние миграций базы данных",
		RunE: func(cmd *cobra.Command, args []string) error {
			db := postgres.NewConnectOrFail(cfg.Postgres)
			mg := migrator.NewOrFail(db)
			status := mg.Status()

			fmt.Printf("\n")
			fmt.Printf("Database migration status\n")
			fmt.Printf("\n")
			fmt.Printf("Version : %d\n", status.Version)
			fmt.Printf("Dirty   : %t\n", status.Dirty)
			fmt.Printf("Error   : %v\n", status.Error)
			fmt.Printf("\n")

			return nil
		},
	}

	rootCmd.AddCommand(cmd)
}

func initSeedMakeCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "seed:make",
		Short: "Создание сидера данных",
		RunE: func(cmd *cobra.Command, args []string) error {
			rowsCount, _ := cmd.Flags().GetInt("rows")
			if rowsCount < 1 {
				fmt.Println("rows must be greater than 0")
				os.Exit(1)
			}

			minAttrs, _ := cmd.Flags().GetInt("min-attrs")
			if minAttrs < 1 || minAttrs > 100 {
				fmt.Println("min-attrs must be in the range from 1 to 100")
				os.Exit(1)
			}

			maxAttrs, _ := cmd.Flags().GetInt("max-attrs")
			if maxAttrs < 1 || maxAttrs > 100 {
				fmt.Println("max-attrs must be in the range from 1 to 100")
				os.Exit(1)
			}

			if minAttrs > maxAttrs {
				fmt.Println("min-attrs must be less than max-attrs")
				os.Exit(1)
			}

			fmt.Printf("\n")
			fmt.Printf("Generating seed csv file with %d rows...\n", rowsCount)

			file, err := seeder.GenerateRows(rowsCount, minAttrs, maxAttrs)
			if err == nil {
				fmt.Printf("Created file %s\n", file)
			} else {
				fmt.Printf("Error: %v", err)
			}
			fmt.Printf("\n")

			return err
		},
	}

	cmd.PersistentFlags().Int("rows", 10, "Number of rows to generate, default 10")
	cmd.PersistentFlags().Int("min-attrs", 10, "Minimum count attributes (1-100), default 10")
	cmd.PersistentFlags().Int("max-attrs", 100, "Maximum count attributes (1-100), default 100")

	rootCmd.AddCommand(cmd)
}

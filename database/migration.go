package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/enricoanto/final-project/config"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			command := args[0]
			cmdArgs := args[1:]

			goose.SetDialect("postgres")

			if command == "create-database" || command == "drop-database" {
				db, err := sql.Open("postgres", config.GetDBConfig().GetDBUrl())
				if err != nil {
					log.Fatal(err)
				}
				switch command {
				case "create-database":
					_, err = db.Exec("CREATE DATABASE " + config.GetDBConfig().DBName)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("success create database: %v", config.GetDBConfig().DBName)
					return
				case "drop-database":
					_, err = db.Exec("DROP DATABASE " + config.GetDBConfig().DBName)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("success drop database: %v", config.GetDBConfig().DBName)
					return

				}
			} else {
				dsn := fmt.Sprintf(config.GetDBConfig().GetDBUrl()+" database=%v", config.GetDBConfig().DBName)
				db, err := sql.Open("postgres", dsn)
				if err != nil {
					log.Fatal(err)
				}
				appPath, _ := os.Getwd()
				dir := appPath + "/database/migration"
				if len(args) == 0 {
					cmd.Help()
					os.Exit(0)
				}

				err = goose.Run(command, db, dir, cmdArgs...)
				if err != nil {
					log.Fatalf("goose run error: %v", err)
				}
			}

		},
	}

	rootCmd.Execute()
}

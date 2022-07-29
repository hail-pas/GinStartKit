package main

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/initialize"
	"os"
)

/*
migrate create -ext sql -dir storage/relational/migrations -seq rbac
*/

func main() {
	configFile := flag.String("conf", "./config/content/default.yaml", "Path to the configuration file")
	migrationSource := flag.String("source", "file://./storage/relational/migrations", "Path to the migration file")
	dsn := flag.String("dsn", "", "database dsn")
	operation := flag.String("op", "", "operation")
	version := flag.Uint("version", 0, "version")

	flag.Parse()

	if *migrationSource == "" {
		fmt.Println("migration source path required")
		os.Exit(0)
	}

	if *operation == "" {
		fmt.Println("operation required")
		os.Exit(0)
	}

	if *dsn == "" {
		initialize.Configuration(*configFile)
		*dsn = global.Configuration.RelationalDatabase.Dsn()
	}

	fmt.Printf("migrations file source: %s\n", *migrationSource)

	m, err := migrate.New(
		*migrationSource,
		*dsn,
	)
	if err != nil {
		panic(err)
	}

	switch *operation {
	case "up":
		err := m.Up()
		if err != nil {
			fmt.Println(err)
		}
	case "down":
		err := m.Down()
		if err != nil {
			fmt.Println(err)
		}
	case "migrate":
		err := m.Migrate(*version)
		if err != nil {
			fmt.Println(err)
		}
	case "force":
		err := m.Force(int(*version))
		if err != nil {
			fmt.Println(err)
		}
	case "version":
		currentVersion, isDirty, err := m.Version()
		if err != nil {
			fmt.Println("err")
			return
		}
		fmt.Printf("currentVersion: %d  %t\n", currentVersion, isDirty)
	}

}

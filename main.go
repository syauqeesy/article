package main

import (
	"log"
	"os"

	"ahmadsyauqi.dev/article/foundation"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run main.go [http|migration]")
	}

	if os.Args[1] == foundation.FoundationMigration && len(os.Args) < 3 {
		log.Fatal("usage: go run main.go [generate|up|down]")
	}

	foundationType := os.Args[1]
	arguments := make([]string, 0, 2)

	if os.Args[1] == foundation.FoundationMigration && os.Args[2] == foundation.MigrationGenerate && len(os.Args) > 3 {
		arguments = append(arguments, os.Args[2], os.Args[3])
	} else if os.Args[1] == foundation.FoundationMigration && os.Args[2] != foundation.MigrationGenerate && len(os.Args) < 3 {
		log.Fatal("usage: go run main.go [generate {migration_name}]")
	} else if os.Args[1] == foundation.FoundationMigration && os.Args[2] != foundation.MigrationGenerate {
		arguments = append(arguments, os.Args[2])
	}

	err := foundation.Boot(foundationType, arguments)
	if err != nil {
		log.Fatal(err.Error())
	}
}

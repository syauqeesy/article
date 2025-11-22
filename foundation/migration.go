package foundation

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"ahmadsyauqi.dev/article/configuration"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type migrationFoundation struct {
	configuration         *configuration.Configuration
	migrator              *migrate.Migrate
	databaseConnectionUrl string
	commandType           string
	commandArgument       string
	migrationPath         string
}

const (
	MigrationGenerate = "generate"
	MigrationUp       = "up"
	MigrationDown     = "down"
)

func (f *migrationFoundation) Setup() (err error) {
	f.databaseConnectionUrl = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s", f.configuration.Database.User, f.configuration.Database.Password, f.configuration.Database.Host, f.configuration.Database.Port, f.configuration.Database.Name, f.configuration.Database.Sslmode)

	f.migrationPath, err = filepath.Abs("migration")
	if err != nil {
		return fmt.Errorf("failed to initialize migration: %w", err)
	}

	if f.commandType != MigrationUp && f.commandType != MigrationDown {
		return nil
	}

	fmt.Println(f.migrationPath)
	m, err := migrate.New("file://"+f.migrationPath, f.databaseConnectionUrl)
	if err != nil {
		return fmt.Errorf("failed to initialize migration: %w", err)
	}

	f.migrator = m

	return nil
}

func (f *migrationFoundation) Boot() error {
	switch f.commandType {
	case MigrationGenerate:
		err := f.Generate(f.commandArgument)
		if err != nil {
			return err
		}

		return nil
	case MigrationUp:
		err := f.migrator.Up()
		if err != nil {
			return err
		}

		return nil
	case MigrationDown:
		err := f.migrator.Down()
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (f *migrationFoundation) Shutdown() error {
	if f.migrator == nil {
		return nil
	}

	sourceErr, databaseErr := f.migrator.Close()
	if sourceErr != nil {
		return sourceErr
	}

	if databaseErr != nil {
		return databaseErr
	}

	return nil
}

func (f *migrationFoundation) Generate(name string) error {
	timestamp := time.Now().Unix()

	up := filepath.Join(f.migrationPath, fmt.Sprintf("%d_%s.up.sql", timestamp, name))
	down := filepath.Join(f.migrationPath, fmt.Sprintf("%d_%s.down.sql", timestamp, name))

	err := os.MkdirAll(f.migrationPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create migrations directory: %w", err)
	}

	for _, path := range []string{up, down} {
		file, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("failed to create %s: %w", path, err)
		}

		file.Close()
	}

	fmt.Println(" -", up)
	fmt.Println(" -", down)

	return nil
}

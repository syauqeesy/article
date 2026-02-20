package foundation

import (
	"fmt"
	"net/url"

	"ahmadsyauqi.dev/article/configuration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type databaseFoundation struct {
	configuration         *configuration.Configuration
	database              *gorm.DB
	databaseConfiguration *gorm.Config
	databaseConnectionUrl string
}

func (f *databaseFoundation) Setup() error {
	f.databaseConnectionUrl = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s&TimeZone=%s&connect_timeout=%d", f.configuration.Database.User, url.QueryEscape(f.configuration.Database.Password), f.configuration.Database.Host, f.configuration.Database.Port, f.configuration.Database.Name, f.configuration.Database.Sslmode, f.configuration.Database.Timezone, f.configuration.Database.Timeout)

	f.databaseConfiguration = &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	return nil
}

func (f *databaseFoundation) Boot() (err error) {
	f.database, err = gorm.Open(postgres.Open(f.databaseConnectionUrl), f.databaseConfiguration)
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %w", err)
	}

	fmt.Println("Connected to database")

	return nil
}

func (f *databaseFoundation) Shutdown() error {
	database, err := f.database.DB()
	if err != nil {
		return fmt.Errorf("Failed to get database object: %w", err)
	}

	err = database.Close()
	if err != nil {
		return fmt.Errorf("Failed to close connection: %w", err)
	}

	fmt.Println("Database connection closed.")

	return nil
}

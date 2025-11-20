package foundation

import (
	"errors"
	"path/filepath"

	"ahmadsyauqi.dev/article/configuration"
)

type Foundation interface {
	Setup() (err error)
	Boot() (err error)
	Shutdown() (err error)
}

const (
	FoundationHttp      = "http"
	FoundationMigration = "migration"
)

func Boot(foundationType string, arguments []string) error {
	foundation, err := New(foundationType, arguments)
	if err != nil {
		return err
	}

	defer foundation.Shutdown()

	err = foundation.Setup()
	if err != nil {
		return err
	}

	err = foundation.Boot()
	if err != nil {
		return err
	}

	return nil
}

func New(foundationType string, arguments []string) (Foundation, error) {
	path, err := filepath.Abs("./config.json")
	if err != nil {
		return nil, err
	}

	configuration, err := configuration.Load(path)
	if err != nil {
		return nil, err
	}

	switch foundationType {
	case FoundationHttp:
		return &httpFoundation{
			configuration: configuration,
		}, nil
	default:
		return nil, errors.New("invalid foundation type")
	}
}

package repository

import "gorm.io/gorm"

type repository struct {
	Database *gorm.DB
}

type Repository struct {
	Account AccountRepository
}

func New(database *gorm.DB) *Repository {
	repository := &repository{
		Database: database,
	}

	return &Repository{
		Account: (*accountRepository)(repository),
	}
}

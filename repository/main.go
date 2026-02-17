package repository

import "gorm.io/gorm"

type repository struct {
	Database *gorm.DB
}

type Repository struct {
	Account           AccountRepository
	AccountIdentity   AccountIdentityRepository
	RefreshToken      RefreshTokenRepository
	Permission        PermissionRepository
	AccountPermission AccountPermissionRepository
	Article           ArticleRepository
	ArticleContent    ArticleContentRepository
}

func New(database *gorm.DB) *Repository {
	repository := &repository{
		Database: database,
	}

	return &Repository{
		Account:           (*accountRepository)(repository),
		AccountIdentity:   (*accountIdentityRepository)(repository),
		RefreshToken:      (*refreshTokenRepository)(repository),
		Permission:        (*permissionRepository)(repository),
		AccountPermission: (*accountPermissionRepository)(repository),
		Article:           (*articleRepository)(repository),
		ArticleContent:    (*articleContentRepository)(repository),
	}
}

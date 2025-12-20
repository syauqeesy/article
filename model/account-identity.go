package model

type AccountIdentity struct {
	Id             string `gorm:"column:id;type:char(36);primaryKey"`
	AccountID      string `gorm:"column:account_id;type:char(36);not null"`
	Provider       string `gorm:"column:provider;type:varchar(32);not null"`
	ProviderUserID string `gorm:"column:provider_user_id;type:varchar(128);not null"`
	CreatedAt      int64  `gorm:"column:created_at;type:bigint;not null"`
	UpdatedAt      *int64 `gorm:"column:updated_at;type:bigint;default:null"`
	DeletedAt      *int64 `gorm:"column:deleted_at;type:bigint;default:null"`

	Account Account `gorm:"foreignKey:AccountID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (AccountIdentity) TableName() string {
	return "account_identities"
}

package model

type Account struct {
	Id        string `gorm:"column:id;type:char(36);primaryKey"`
	Email     string `gorm:"column:email;varchar(128);not null"`
	CreatedAt int64  `gorm:"column:created_at;type:bigint;not null"`
	UpdatedAt *int64 `gorm:"column:updated_at;type:bigint;default:null"`
	DeletedAt *int64 `gorm:"column:deleted_at;type:bigint;default:null"`
}

func (Account) TableName() string {
	return "accounts"
}

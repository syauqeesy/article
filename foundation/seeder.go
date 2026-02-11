package foundation

import (
	"log"
	"time"

	"ahmadsyauqi.dev/article/configuration"
	"ahmadsyauqi.dev/article/model"
)

type seederFoundation struct {
	configuration *configuration.Configuration
	database      *databaseFoundation
}

func (f *seederFoundation) Setup() error {
	f.database = &databaseFoundation{
		configuration: f.configuration,
	}

	err := f.database.Setup()
	if err != nil {
		return err
	}

	err = f.database.Boot()
	if err != nil {
		return err
	}

	return nil
}

func (f *seederFoundation) Boot() error {
	var t time.Time
	result := f.database.database.Raw("SELECT NOW() AS THIS_MOMENT").Scan(&t)
	if result.Error != nil {
		log.Fatal("database error : ", result.Error)
	}

	var err error

	account := &model.Account{}

	q := f.database.database.Where("email = ?", "syauqeesy@gmail.com").First(&account)
	if q.Error != nil {
		account, err = model.NewAccount("syauqeesy@gmail.com", "Ahmad Syauqi")
		if err != nil {
			return err
		}

		q = f.database.database.Create(&account)
		if q.Error != nil {
			return q.Error
		}

		accountIdentity, err := model.NewAccountIdentity(account.GetId(), "google", "115511328782817718638")
		if err != nil {
			return err
		}

		q = f.database.database.Create(&accountIdentity)
		if q.Error != nil {
			return q.Error
		}
	}

	permissions := make([]*model.Permission, 0)

	values := [][]string{
		{"article.show", "Show Article"},
		{"article.list", "List Article"},
		{"article.create", "Create Article"},
		{"article.update", "Update Article"},
		{"article.publish", "Publish Article"},
	}

	for _, value := range values {
		permission := &model.Permission{}

		q = f.database.database.Where("code = ?", value[0]).First(&permission)
		if q.Error != nil {
			permission, err = model.NewPermission(value[0], value[1])
			if err != nil {
				return err
			}

			q = f.database.database.Create(&permission)
			if q.Error != nil {
				return q.Error
			}
		}

		permissions = append(permissions, permission)
	}

	for _, permission := range permissions {
		accountPermission := &model.AccountPermission{}

		q := f.database.database.Where("account_id = ?", account.GetId()).Where("permission_id = ?", permission.GetId()).First(&accountPermission)
		if q.Error != nil {
			accountPermission, err = model.NewAccountPermission(account.GetId(), permission.GetId())
			if err != nil {
				return err
			}

			q = f.database.database.Create(&accountPermission)
			if q.Error != nil {
				return q.Error
			}
		}
	}

	return nil
}

func (f *seederFoundation) Shutdown() error {
	err := f.database.Shutdown()
	if err != nil {
		return err
	}

	return nil
}

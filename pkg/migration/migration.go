package migration

import (
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, getMigrations())

	if err := m.Migrate(); err != nil {
		return err
	}
	return nil
}

func readSQLFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func getMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "20240212_create_users_table",
			Migrate: func(tx *gorm.DB) error {
				sql, err := readSQLFile("pkg/migration/20240212_create_users_table.up.sql")
				if err != nil {
					return err
				}
				return tx.Exec(sql).Error
			},
			Rollback: func(tx *gorm.DB) error {
				sql, err := readSQLFile("pkg/migration/20240212_create_users_table.down.sql")
				if err != nil {
					return err
				}
				return tx.Exec(sql).Error
			},
		},
	}
}

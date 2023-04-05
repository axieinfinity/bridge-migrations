package models

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20230504",
		Migrate: func(tx *gorm.DB) error {
			// drop all tables
			if err := tx.Exec("truncate table deposit").Error; err != nil {
				return err
			}
			if err := tx.Exec("truncate table event").Error; err != nil {
				return err
			}
			if err := tx.Exec("truncate table job").Error; err != nil {
				return err
			}
			if err := tx.Exec("truncate table processed_block").Error; err != nil {
				return err
			}
			if err := tx.Exec("truncate table processed_receipt").Error; err != nil {
				return err
			}
			if err := tx.Exec("truncate table task").Error; err != nil {
				return err
			}
			if err := tx.Exec("truncate table withdrawal").Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}
}

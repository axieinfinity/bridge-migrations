package migration

import (
	"strings"

	bridgeCore "github.com/axieinfinity/bridge-core"
	models20220515 "github.com/axieinfinity/bridge-migrations/20220515_init"
	models20220701 "github.com/axieinfinity/bridge-migrations/20220701_retry_all_failed_tx"
	models20220703 "github.com/axieinfinity/bridge-migrations/20220703_add_timestamp_when_submit_tx"
	models20220707 "github.com/axieinfinity/bridge-migrations/20220707_add_processed_receipt_table"
	models20220708 "github.com/axieinfinity/bridge-migrations/20220708_update_processed_block"
	models20230504 "github.com/axieinfinity/bridge-migrations/20230504_drop_old_data"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, cfg *bridgeCore.Config) error {
	migrations := []*gormigrate.Migration{
		models20220515.Migrate(),
		models20220701.Migrate(),
		models20220703.Migrate(),
		models20220707.Migrate(),
		models20220708.Migrate(),
		models20230504.Migrate(),
	}
	migrate := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	for _, version := range migrations {
		if err := migrate.MigrateTo(version.ID); err != nil {
			if strings.Contains(err.Error(), "Duplicated") {
				continue
			}
			return err
		}
	}
	return nil
}

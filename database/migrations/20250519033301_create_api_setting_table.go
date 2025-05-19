package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519033301CreateApiSettingTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519033301CreateApiSettingTable) Signature() string {
	return "20250519033301_create_api_setting_table"
}

// Up Run the migrations.
func (r *M20250519033301CreateApiSettingTable) Up() error {
	if !facades.Schema().HasTable("api_setting") {
		return facades.Schema().Create("api_setting", func(table schema.Blueprint) {
			table.ID("id")
			table.Integer("auto_sync").Default(0).Nullable()
			table.Integer("edit_profile_siswa").Default(0).Nullable()
			table.Integer("edit_profile_guru").Default(0).Nullable()
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519033301CreateApiSettingTable) Down() error {
 	return facades.Schema().DropIfExists("api_setting")
}

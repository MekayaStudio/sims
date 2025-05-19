package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519042551CreateLevelKelasTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519042551CreateLevelKelasTable) Signature() string {
	return "20250519042551_create_level_kelas_table"
}

// Up Run the migrations.
func (r *M20250519042551CreateLevelKelasTable) Up() error {
	if !facades.Schema().HasTable("level_kelas") {
		return facades.Schema().Create("level_kelas", func(table schema.Blueprint) {
			table.Integer("id_level")
			table.String("level", 2).Nullable()
			table.Index("id_level")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519042551CreateLevelKelasTable) Down() error {
 	return facades.Schema().DropIfExists("level_kelas")
}

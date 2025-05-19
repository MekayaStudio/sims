package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519042544CreateLevelGuruTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519042544CreateLevelGuruTable) Signature() string {
	return "20250519042544_create_level_guru_table"
}

// Up Run the migrations.
func (r *M20250519042544CreateLevelGuruTable) Up() error {
	if !facades.Schema().HasTable("level_guru") {
		return facades.Schema().Create("level_guru", func(table schema.Blueprint) {
			table.Increments("id_level")
			table.String("level", 50).Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519042544CreateLevelGuruTable) Down() error {
 	return facades.Schema().DropIfExists("level_guru")
}

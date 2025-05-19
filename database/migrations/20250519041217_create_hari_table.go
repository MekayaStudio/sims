package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519041217CreateHariTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519041217CreateHariTable) Signature() string {
	return "20250519041217_create_hari_table"
}

// Up Run the migrations.
func (r *M20250519041217CreateHariTable) Up() error {
	if !facades.Schema().HasTable("hari") {
		return facades.Schema().Create("hari", func(table schema.Blueprint) {
			table.Increments("id_hri")
			table.String("nama_hri", 50)
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519041217CreateHariTable) Down() error {
 	return facades.Schema().DropIfExists("hari")
}

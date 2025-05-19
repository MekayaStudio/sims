package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519033310CreateBulanTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519033310CreateBulanTable) Signature() string {
	return "20250519033310_create_bulan_table"
}

// Up Run the migrations.
func (r *M20250519033310CreateBulanTable) Up() error {
	if !facades.Schema().HasTable("bulan") {
		return facades.Schema().Create("bulan", func(table schema.Blueprint) {
			table.ID("id_bln")
			table.String("nama_bln", 25)
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519033310CreateBulanTable) Down() error {
 	return facades.Schema().DropIfExists("bulan")
}

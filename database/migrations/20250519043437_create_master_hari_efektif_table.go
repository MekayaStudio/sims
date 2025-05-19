package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519043437CreateMasterHariEfektifTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519043437CreateMasterHariEfektifTable) Signature() string {
	return "20250519043437_create_master_hari_efektif_table"
}

// Up Run the migrations.
func (r *M20250519043437CreateMasterHariEfektifTable) Up() error {
	if !facades.Schema().HasTable("master_hari_efektif") {
		return facades.Schema().Create("master_hari_efektif", func(table schema.Blueprint) {
			table.Increments("id_hari_efektif")
			table.Integer("jml_hari_efektif")
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519043437CreateMasterHariEfektifTable) Down() error {
 	return facades.Schema().DropIfExists("master_hari_efektif")
}

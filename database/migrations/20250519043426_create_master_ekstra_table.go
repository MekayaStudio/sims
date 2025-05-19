package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519043426CreateMasterEkstraTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519043426CreateMasterEkstraTable) Signature() string {
	return "20250519043426_create_master_ekstra_table"
}

// Up Run the migrations.
func (r *M20250519043426CreateMasterEkstraTable) Up() error {
	if !facades.Schema().HasTable("master_ekstra") {
		return facades.Schema().Create("master_ekstra", func(table schema.Blueprint) {
			table.Increments("id_ekstra")
			table.String("nama_ekstra", 100)
			table.String("kode_ekstra", 20)
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519043426CreateMasterEkstraTable) Down() error {
 	return facades.Schema().DropIfExists("master_ekstra")
}

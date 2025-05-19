package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519045026CreateRaporNaikTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519045026CreateRaporNaikTable) Signature() string {
	return "20250519045026_create_rapor_naik_table"
}

// Up Run the migrations.
func (r *M20250519045026CreateRaporNaikTable) Up() error {
	if !facades.Schema().HasTable("rapor_naik") {
		return facades.Schema().Create("rapor_naik", func(table schema.Blueprint) {
			table.Integer("id_naik")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("id_siswa")
			table.Integer("naik")
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519045026CreateRaporNaikTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_naik")
}

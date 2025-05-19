package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519044502CreateRaporDataFisikTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519044502CreateRaporDataFisikTable) Signature() string {
	return "20250519044502_create_rapor_data_fisik_table"
}

// Up Run the migrations.
func (r *M20250519044502CreateRaporDataFisikTable) Up() error {
	if !facades.Schema().HasTable("rapor_data_fisik") {
		return facades.Schema().Create("rapor_data_fisik", func(table schema.Blueprint) {
			table.Increments("id_fisik")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("id_kelas").Nullable()
			table.Integer("jenis")
			table.Integer("kode")
			table.LongText("deskripsi")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519044502CreateRaporDataFisikTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_data_fisik")
}

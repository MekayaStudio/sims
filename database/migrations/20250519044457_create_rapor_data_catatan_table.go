package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519044457CreateRaporDataCatatanTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519044457CreateRaporDataCatatanTable) Signature() string {
	return "20250519044457_create_rapor_data_catatan_table"
}

// Up Run the migrations.
func (r *M20250519044457CreateRaporDataCatatanTable) Up() error {
	if !facades.Schema().HasTable("rapor_data_catatan") {
		return facades.Schema().Create("rapor_data_catatan", func(table schema.Blueprint) {
			table.Increments("id_catatan")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("id_kelas").Nullable()
			table.Integer("jenis")
			table.Integer("kode")
			table.String("deskripsi", 150)
			table.String("rank", 7).Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519044457CreateRaporDataCatatanTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_data_catatan")
}

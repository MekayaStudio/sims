package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519044513CreateRaporFisikTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519044513CreateRaporFisikTable) Signature() string {
	return "20250519044513_create_rapor_fisik_table"
}

// Up Run the migrations.
func (r *M20250519044513CreateRaporFisikTable) Up() error {
	if !facades.Schema().HasTable("rapor_fisik") {
		return facades.Schema().Create("rapor_fisik", func(table schema.Blueprint) {
			table.Increments("id_fisik")
			table.Integer("id_kelas").Nullable()
			table.Integer("id_siswa").Nullable()
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.LongText("kondisi")
			table.Integer("tinggi")
			table.Integer("berat")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519044513CreateRaporFisikTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_fisik")
}

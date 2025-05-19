package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519044452CreateRaporCatatanWaliTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519044452CreateRaporCatatanWaliTable) Signature() string {
	return "20250519044452_create_rapor_catatan_wali_table"
}

// Up Run the migrations.
func (r *M20250519044452CreateRaporCatatanWaliTable) Up() error {
	if !facades.Schema().HasTable("rapor_catatan_wali") {
		return facades.Schema().Create("rapor_catatan_wali", func(table schema.Blueprint) {
			table.Increments("id_catatan_wali")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("id_kelas").Nullable()
			table.Integer("id_siswa").Nullable()
			table.LongText("nilai").Nullable()
			table.LongText("deskripsi").Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519044452CreateRaporCatatanWaliTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_catatan_wali")
}

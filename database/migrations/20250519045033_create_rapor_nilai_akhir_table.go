package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519045033CreateRaporNilaiAkhirTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519045033CreateRaporNilaiAkhirTable) Signature() string {
	return "20250519045033_create_rapor_nilai_akhir_table"
}

// Up Run the migrations.
func (r *M20250519045033CreateRaporNilaiAkhirTable) Up() error {
	if !facades.Schema().HasTable("rapor_nilai_akhir") {
		return facades.Schema().Create("rapor_nilai_akhir", func(table schema.Blueprint) {
			table.Increments("id_nilai_akhir")
			table.Integer("id_mapel").Nullable()
			table.Integer("id_kelas").Nullable()
			table.Integer("id_siswa").Nullable()
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("nilai").Default(0)
			table.Integer("akhir").Nullable()
			table.String("predikat", 1).Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519045033CreateRaporNilaiAkhirTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_nilai_akhir")
}

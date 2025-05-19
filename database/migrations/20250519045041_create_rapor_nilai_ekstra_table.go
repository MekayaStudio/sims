package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519045041CreateRaporNilaiEkstraTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519045041CreateRaporNilaiEkstraTable) Signature() string {
	return "20250519045041_create_rapor_nilai_ekstra_table"
}

// Up Run the migrations.
func (r *M20250519045041CreateRaporNilaiEkstraTable) Up() error {
	if !facades.Schema().HasTable("rapor_nilai_ekstra") {
		return facades.Schema().Create("rapor_nilai_ekstra", func(table schema.Blueprint) {
			table.Increments("id_nilai_ekstra")
			table.Integer("id_ekstra").Nullable()
			table.Integer("id_kelas").Nullable()
			table.Integer("id_siswa").Nullable()
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("nilai")
			table.String("predikat", 1).Nullable()
			table.Text("deskripsi").Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519045041CreateRaporNilaiEkstraTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_nilai_ekstra")
}

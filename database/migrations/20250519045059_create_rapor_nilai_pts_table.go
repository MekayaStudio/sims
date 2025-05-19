package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519045059CreateRaporNilaiPtsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519045059CreateRaporNilaiPtsTable) Signature() string {
	return "20250519045059_create_rapor_nilai_pts_table"
}

// Up Run the migrations.
func (r *M20250519045059CreateRaporNilaiPtsTable) Up() error {
	if !facades.Schema().HasTable("rapor_nilai_pts") {
		return facades.Schema().Create("rapor_nilai_pts", func(table schema.Blueprint) {
			table.Increments("id_nilai_pts")
			table.Integer("id_mapel").Nullable()
			table.Integer("id_kelas").Nullable()
			table.Integer("id_siswa").Nullable()
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("nilai").Default(0)
			table.String("predikat", 1).Nullable()
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519045059CreateRaporNilaiPtsTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_nilai_pts")
}

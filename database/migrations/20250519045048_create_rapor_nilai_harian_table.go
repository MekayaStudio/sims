package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519045048CreateRaporNilaiHarianTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519045048CreateRaporNilaiHarianTable) Signature() string {
	return "20250519045048_create_rapor_nilai_harian_table"
}

// Up Run the migrations.
func (r *M20250519045048CreateRaporNilaiHarianTable) Up() error {
	if !facades.Schema().HasTable("rapor_nilai_harian") {
		return facades.Schema().Create("rapor_nilai_harian", func(table schema.Blueprint) {
			table.Increments("id_nilai_harian")
			table.Integer("id_siswa").Nullable()
			table.Integer("id_mapel").Nullable()
			table.Integer("id_kelas").Nullable()
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.String("p1", 3).Nullable()
			table.String("p2", 3).Nullable()
			table.String("p3", 3).Nullable()
			table.String("p4", 3).Nullable()
			table.String("p5", 3).Nullable()
			table.String("p6", 3).Nullable()
			table.String("p7", 3).Nullable()
			table.String("p8", 3).Nullable()
			table.String("p_rata_rata", 4).Nullable()
			table.String("p_predikat", 1).Nullable()
			table.LongText("p_deskripsi").Nullable()
			table.String("k1", 3).Nullable()
			table.String("k2", 3).Nullable()
			table.String("k3", 3).Nullable()
			table.String("k4", 3).Nullable()
			table.String("k5", 3).Nullable()
			table.String("k6", 3).Nullable()
			table.String("k7", 3).Nullable()
			table.String("k8", 3).Nullable()
			table.String("k_rata_rata", 4).Nullable()
			table.String("k_predikat", 1).Nullable()
			table.LongText("k_deskripsi").Nullable()
			table.Integer("jml").Default(0)
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519045048CreateRaporNilaiHarianTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_nilai_harian")
}

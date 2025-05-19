package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519042457CreateKelasJadwalMateriTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519042457CreateKelasJadwalMateriTable) Signature() string {
	return "20250519042457_create_kelas_jadwal_materi_table"
}

// Up Run the migrations.
func (r *M20250519042457CreateKelasJadwalMateriTable) Up() error {
	if !facades.Schema().HasTable("kelas_jadwal_materi") {
		return facades.Schema().Create("kelas_jadwal_materi", func(table schema.Blueprint) {
			table.String("id_kjm", 50)
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("id_materi").Nullable()
			table.Integer("id_mapel").Nullable()
			table.Integer("id_kelas").Nullable()
			table.String("jadwal_materi", 20)
			table.Integer("jenis").Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519042457CreateKelasJadwalMateriTable) Down() error {
	return facades.Schema().DropIfExists("kelas_jadwal_materi")
}

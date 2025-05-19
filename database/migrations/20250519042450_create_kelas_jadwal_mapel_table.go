package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519042450CreateKelasJadwalMapelTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519042450CreateKelasJadwalMapelTable) Signature() string {
	return "20250519042450_create_kelas_jadwal_mapel_table"
}

// Up Run the migrations.
func (r *M20250519042450CreateKelasJadwalMapelTable) Up() error {
	if !facades.Schema().HasTable("kelas_jadwal_mapel") {
		return facades.Schema().Create("kelas_jadwal_mapel", func(table schema.Blueprint) {
			table.Integer("id_jadwal")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("id_kelas").Nullable()
			table.Integer("id_hari")
			table.Integer("jam_ke")
			table.Integer("id_mapel").Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519042450CreateKelasJadwalMapelTable) Down() error {
 	return facades.Schema().DropIfExists("kelas_jadwal_mapel")
}

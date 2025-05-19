package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519041255CreateKelasJadwalKbmTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519041255CreateKelasJadwalKbmTable) Signature() string {
	return "20250519041255_create_kelas_jadwal_kbm_table"
}

// Up Run the migrations.
func (r *M20250519041255CreateKelasJadwalKbmTable) Up() error {
	if !facades.Schema().HasTable("kelas_jadwal_kbm") {
		return facades.Schema().Create("kelas_jadwal_kbm", func(table schema.Blueprint) {
			table.Integer("id_kbm")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("id_kelas").Nullable()
			table.Integer("kbm_jam_pel")
			table.String("kbm_jam_mulai", 5)
			table.Integer("kbm_jml_mapel_hari")
			table.Text("istirahat")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519041255CreateKelasJadwalKbmTable) Down() error {
 	return facades.Schema().DropIfExists("kelas_jadwal_kbm")
}

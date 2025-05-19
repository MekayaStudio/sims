package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519033319CreateCbtJadwalTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519033319CreateCbtJadwalTable) Signature() string {
	return "20250519033319_create_cbt_jadwal_table"
}

// Up Run the migrations.
func (r *M20250519033319CreateCbtJadwalTable) Up() error {
	if !facades.Schema().HasTable("cbt_jadwal") {
		return facades.Schema().Create("cbt_jadwal", func(table schema.Blueprint) {
			table.ID("id_jadwal")
			table.Char("id_tp", 2)
			table.Char("id_smt", 2)
			table.Integer("id_bank")
			table.Integer("id_jenis")
			table.String("tgl_mulai", 20)
			table.String("tgl_selesai", 20)
			table.Integer("durasi_ujian")
			table.LongText("pengawas")
			table.Integer("acak_soal")
			table.Integer("acak_opsi")
			table.Integer("hasil_tampil")
			table.Integer("token")
			table.Integer("status")
			table.Integer("ulang")
			table.Integer("reset_login")
			table.Integer("rekap").Default(0)
			table.Integer("jam_ke").Default(0)
			table.Integer("jarak").Default(0)
			table.DateTime("time_create")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519033319CreateCbtJadwalTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_jadwal")
}

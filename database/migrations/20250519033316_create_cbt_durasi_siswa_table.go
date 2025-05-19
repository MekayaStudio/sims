package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519033316CreateCbtDurasiSiswaTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519033316CreateCbtDurasiSiswaTable) Signature() string {
	return "20250519033316_create_cbt_durasi_siswa_table"
}

// Up Run the migrations.
func (r *M20250519033316CreateCbtDurasiSiswaTable) Up() error {
	if !facades.Schema().HasTable("cbt_durasi_siswa") {
		return facades.Schema().Create("cbt_durasi_siswa", func(table schema.Blueprint) {
			table.Integer("id_durasi")
			table.Integer("id_siswa")
			table.Integer("id_jadwal")
			table.Integer("status").Default(0)
			table.Time("lama_ujian")
			table.String("mulai", 22)
			table.String("selesai", 22)
			table.Integer("reset").Default(0)
			table.DateTime("time_create")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519033316CreateCbtDurasiSiswaTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_durasi_siswa")
}

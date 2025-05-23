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
			table.ID("id_durasi")
			table.Integer("id_siswa").Nullable()
			table.Integer("id_jadwal").Nullable()
			table.Integer("status").Default(0).Nullable()
			table.Time("lama_ujian").Nullable()
			table.String("mulai", 22).Nullable()
			table.String("selesai", 22).Nullable()
			table.Integer("reset").Default(0).Nullable()
			table.DateTime("time_create").UseCurrent()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519033316CreateCbtDurasiSiswaTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_durasi_siswa")
}

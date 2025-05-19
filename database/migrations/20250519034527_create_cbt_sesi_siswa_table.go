package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519034527CreateCbtSesiSiswaTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519034527CreateCbtSesiSiswaTable) Signature() string {
	return "20250519034527_create_cbt_sesi_siswa_table"
}

// Up Run the migrations.
func (r *M20250519034527CreateCbtSesiSiswaTable) Up() error {
	if !facades.Schema().HasTable("cbt_sesi_siswa") {
		return facades.Schema().Create("cbt_sesi_siswa", func(table schema.Blueprint) {
			table.Integer("siswa_id")
			table.Integer("kelas_id").Nullable()
			table.Integer("ruang_id")
			table.Integer("sesi_id")
			table.Integer("tp_id")
			table.Integer("smt_id")
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519034527CreateCbtSesiSiswaTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_sesi_siswa")
}

package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519042508CreateKelasSiswaTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519042508CreateKelasSiswaTable) Signature() string {
	return "20250519042508_create_kelas_siswa_table"
}

// Up Run the migrations.
func (r *M20250519042508CreateKelasSiswaTable) Up() error {
	if !facades.Schema().HasTable("kelas_siswa") {
		return facades.Schema().Create("kelas_siswa", func(table schema.Blueprint) {
			table.Integer("id_kelas_siswa")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("id_siswa").Nullable()
			table.Integer("id_kelas").Nullable()
			table.Unique("id_kelas_siswa")
			table.Index("id_siswa")
			table.Index("id_kelas")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519042508CreateKelasSiswaTable) Down() error {
 	return facades.Schema().DropIfExists("kelas_siswa")
}

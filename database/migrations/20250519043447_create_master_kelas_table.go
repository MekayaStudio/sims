package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519043447CreateMasterKelasTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519043447CreateMasterKelasTable) Signature() string {
	return "20250519043447_create_master_kelas_table"
}

// Up Run the migrations.
func (r *M20250519043447CreateMasterKelasTable) Up() error {
	if !facades.Schema().HasTable("master_kelas") {
		return facades.Schema().Create("master_kelas", func(table schema.Blueprint) {
			table.Increments("id_kelas")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.String("nama_kelas", 30)
			table.String("kode_kelas", 20)
			table.Integer("jurusan_id")
			table.Integer("level_id")
			table.Integer("guru_id")
			table.Integer("siswa_id")
			table.LongText("jumlah_siswa")
			table.Char("set_siswa", 1).Default("0")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519043447CreateMasterKelasTable) Down() error {
 	return facades.Schema().DropIfExists("master_kelas")
}

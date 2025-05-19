package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519041240CreateKelasCatatanWaliTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519041240CreateKelasCatatanWaliTable) Signature() string {
	return "20250519041240_create_kelas_catatan_wali_table"
}

// Up Run the migrations.
func (r *M20250519041240CreateKelasCatatanWaliTable) Up() error {
	if !facades.Schema().HasTable("kelas_catatan_wali") {
		return facades.Schema().Create("kelas_catatan_wali", func(table schema.Blueprint) {
			table.Increments("id_catatan")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("type").Comment("1=semua siswa, 2=per siswa")
			table.String("level", 1).Comment("1=saran, 2=teguran, 3=peringatan, 4=sangsi")
			table.DateTime("tgl")
			table.Integer("id_siswa").Nullable()
			table.Integer("id_kelas").Nullable()
			table.Text("text")
			table.String("readed", 22).Default("0")
			table.LongText("reading").Nullable()
			table.Integer("jml").Default(0)
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519041240CreateKelasCatatanWaliTable) Down() error {
 	return facades.Schema().DropIfExists("kelas_catatan_wali")
}

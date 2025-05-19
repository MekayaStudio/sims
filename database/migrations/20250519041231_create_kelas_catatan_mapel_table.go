package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519041231CreateKelasCatatanMapelTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519041231CreateKelasCatatanMapelTable) Signature() string {
	return "20250519041231_create_kelas_catatan_mapel_table"
}

// Up Run the migrations.
func (r *M20250519041231CreateKelasCatatanMapelTable) Up() error {
	if !facades.Schema().HasTable("kelas_catatan_mapel") {
		return facades.Schema().Create("kelas_catatan_mapel", func(table schema.Blueprint) {
			table.Increments("id_catatan")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("type")
			table.Integer("id_siswa").Nullable()
			table.Integer("id_mapel").Nullable()
			table.Integer("id_kelas").Nullable()
			table.Integer("id_guru").Nullable()
			table.String("level", 1).Default("0")
			table.DateTime("tgl")
			table.Text("text")
			table.String("readed", 22).Default("0")
			table.LongText("reading").Nullable()
			table.Integer("jml").Default(0)
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519041231CreateKelasCatatanMapelTable) Down() error {
 	return facades.Schema().DropIfExists("kelas_catatan_mapel")
}

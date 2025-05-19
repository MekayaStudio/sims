package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519042503CreateKelasMateriTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519042503CreateKelasMateriTable) Signature() string {
	return "20250519042503_create_kelas_materi_table"
}

// Up Run the migrations.
func (r *M20250519042503CreateKelasMateriTable) Up() error {
	if !facades.Schema().HasTable("kelas_materi") {
		return facades.Schema().Create("kelas_materi", func(table schema.Blueprint) {
			table.Increments("id_materi")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Text("kode_materi")
			table.Integer("id_guru").Nullable()
			table.Text("materi_kelas")
			table.Integer("id_mapel")
			table.String("kode_mapel", 300).Nullable()
			table.Text("judul_materi")
			table.LongText("isi_materi")
			table.LongText("file").Nullable()
			table.String("link_file", 255).Nullable()
			table.DateTime("tgl_mulai").Nullable()
			table.DateTime("created_on").Nullable()
			table.DateTime("updated_on").Nullable()
			table.Integer("status").Nullable()
			table.String("youtube", 255)
			table.Integer("jenis")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519042503CreateKelasMateriTable) Down() error {
 	return facades.Schema().DropIfExists("kelas_materi")
}

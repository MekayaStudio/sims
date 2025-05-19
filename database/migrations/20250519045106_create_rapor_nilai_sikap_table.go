package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519045106CreateRaporNilaiSikapTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519045106CreateRaporNilaiSikapTable) Signature() string {
	return "20250519045106_create_rapor_nilai_sikap_table"
}

// Up Run the migrations.
func (r *M20250519045106CreateRaporNilaiSikapTable) Up() error {
	if !facades.Schema().HasTable("rapor_nilai_sikap") {
		return facades.Schema().Create("rapor_nilai_sikap", func(table schema.Blueprint) {
			table.Increments("id_nilai_sikap")
			table.Integer("id_siswa").Nullable()
			table.Integer("id_kelas").Nullable()
			table.Integer("id_tp").Default(0)
			table.Integer("id_smt").Default(0)
			table.Integer("jenis").Nullable()
			table.LongText("nilai")
			table.LongText("deskripsi").Nullable()
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519045106CreateRaporNilaiSikapTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_nilai_sikap")
}

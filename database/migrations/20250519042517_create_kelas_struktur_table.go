package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519042517CreateKelasStrukturTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519042517CreateKelasStrukturTable) Signature() string {
	return "20250519042517_create_kelas_struktur_table"
}

// Up Run the migrations.
func (r *M20250519042517CreateKelasStrukturTable) Up() error {
	if !facades.Schema().HasTable("kelas_struktur") {
		return facades.Schema().Create("kelas_struktur", func(table schema.Blueprint) {
			table.Increments("id_kelas")
			table.Integer("ketua").Nullable()
			table.Integer("wakil_ketua").Nullable()
			table.Integer("sekretaris_1").Nullable()
			table.Integer("sekretaris_2").Nullable()
			table.Integer("bendahara_1").Nullable()
			table.Integer("bendahara_2").Nullable()
			table.Integer("sie_ekstrakurikuler").Nullable()
			table.Integer("sie_upacara").Nullable()
			table.Integer("sie_olahraga").Nullable()
			table.Integer("sie_keagamaan").Nullable()
			table.Integer("sie_keamanan").Nullable()
			table.Integer("sie_ketertiban").Nullable()
			table.Integer("sie_kebersihan").Nullable()
			table.Integer("sie_keindahan").Nullable()
			table.Integer("sie_kesehatan").Nullable()
			table.Integer("sie_kekeluargaan").Nullable()
			table.Integer("sie_humas").Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519042517CreateKelasStrukturTable) Down() error {
 	return facades.Schema().DropIfExists("kelas_struktur")
}

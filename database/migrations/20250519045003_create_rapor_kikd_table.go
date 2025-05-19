package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519045003CreateRaporKikdTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519045003CreateRaporKikdTable) Signature() string {
	return "20250519045003_create_rapor_kikd_table"
}

// Up Run the migrations.
func (r *M20250519045003CreateRaporKikdTable) Up() error {
	if !facades.Schema().HasTable("rapor_kikd") {
		return facades.Schema().Create("rapor_kikd", func(table schema.Blueprint) {
			table.Increments("id_kikd")
			table.Integer("id_mapel_kelas").Nullable()
			table.Integer("aspek")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.LongText("materi_kikd")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519045003CreateRaporKikdTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_kikd")
}

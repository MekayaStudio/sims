package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519043452CreateMasterKelompokMapelTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519043452CreateMasterKelompokMapelTable) Signature() string {
	return "20250519043452_create_master_kelompok_mapel_table"
}

// Up Run the migrations.
func (r *M20250519043452CreateMasterKelompokMapelTable) Up() error {
	if !facades.Schema().HasTable("master_kelompok_mapel") {
		return facades.Schema().Create("master_kelompok_mapel", func(table schema.Blueprint) {
			table.Increments("id_kel_mapel")
			table.String("kode_kel_mapel", 10)
			table.String("nama_kel_mapel", 100)
			table.String("kategori", 20)
			table.Integer("id_parent")
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519043452CreateMasterKelompokMapelTable) Down() error {
 	return facades.Schema().DropIfExists("master_kelompok_mapel")
}

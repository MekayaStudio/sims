package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519043457CreateMasterMapelTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519043457CreateMasterMapelTable) Signature() string {
	return "20250519043457_create_master_mapel_table"
}

// Up Run the migrations.
func (r *M20250519043457CreateMasterMapelTable) Up() error {
	if !facades.Schema().HasTable("master_mapel") {
		return facades.Schema().Create("master_mapel", func(table schema.Blueprint) {
			table.Increments("id_mapel")
			table.String("nama_mapel", 50)
			table.String("kode", 20)
			table.String("kelompok", 5).Default("-")
			table.Integer("bobot_p").Default(0)
			table.Integer("bobot_k").Default(0)
			table.Integer("jenjang").Default(0)
			table.Integer("urutan")
			table.Integer("status").Default(1)
			table.Integer("deletable").Default(1)
			table.Integer("urutan_tampil")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519043457CreateMasterMapelTable) Down() error {
 	return facades.Schema().DropIfExists("master_mapel")
}

package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519043442CreateMasterJurusanTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519043442CreateMasterJurusanTable) Signature() string {
	return "20250519043442_create_master_jurusan_table"
}

// Up Run the migrations.
func (r *M20250519043442CreateMasterJurusanTable) Up() error {
	if !facades.Schema().HasTable("master_jurusan") {
		return facades.Schema().Create("master_jurusan", func(table schema.Blueprint) {
			table.Increments("id_jurusan")
			table.String("nama_jurusan", 30)
			table.String("kode_jurusan", 10)
			table.LongText("mapel_peminatan")
			table.Integer("status").Default(1)
			table.Integer("deletable").Default(1)
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519043442CreateMasterJurusanTable) Down() error {
 	return facades.Schema().DropIfExists("master_jurusan")
}

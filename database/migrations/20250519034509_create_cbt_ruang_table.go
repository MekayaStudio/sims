package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519034509CreateCbtRuangTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519034509CreateCbtRuangTable) Signature() string {
	return "20250519034509_create_cbt_ruang_table"
}

// Up Run the migrations.
func (r *M20250519034509CreateCbtRuangTable) Up() error {
	if !facades.Schema().HasTable("cbt_ruang") {
		return facades.Schema().Create("cbt_ruang", func(table schema.Blueprint) {
			table.Increments("id_ruang")
			table.String("nama_ruang", 50)
			table.String("kode_ruang", 10)
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519034509CreateCbtRuangTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_ruang")
}

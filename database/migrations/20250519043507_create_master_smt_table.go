package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519043507CreateMasterSmtTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519043507CreateMasterSmtTable) Signature() string {
	return "20250519043507_create_master_smt_table"
}

// Up Run the migrations.
func (r *M20250519043507CreateMasterSmtTable) Up() error {
	if !facades.Schema().HasTable("master_smt") {
		return facades.Schema().Create("master_smt", func(table schema.Blueprint) {
			table.Increments("id_smt")
			table.String("smt", 10)
			table.String("nama_smt", 10)
			table.Integer("active")
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519043507CreateMasterSmtTable) Down() error {
 	return facades.Schema().DropIfExists("master_smt")
}

package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519044420CreateMasterTpTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519044420CreateMasterTpTable) Signature() string {
	return "20250519044420_create_master_tp_table"
}

// Up Run the migrations.
func (r *M20250519044420CreateMasterTpTable) Up() error {
	if !facades.Schema().HasTable("master_tp") {
		return facades.Schema().Create("master_tp", func(table schema.Blueprint) {
			table.Increments("id_tp")
			table.String("tahun", 20)
			table.Integer("active")
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519044420CreateMasterTpTable) Down() error {
 	return facades.Schema().DropIfExists("master_tp")
}

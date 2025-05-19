package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519042624CreateLogTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519042624CreateLogTable) Signature() string {
	return "20250519042624_create_log_table"
}

// Up Run the migrations.
func (r *M20250519042624CreateLogTable) Up() error {
	if !facades.Schema().HasTable("log") {
		return facades.Schema().Create("log", func(table schema.Blueprint) {
			table.Increments("id_log")
			table.DateTime("log_time")
			table.Integer("id_user")
			table.Integer("id_group")
			table.Text("name_group")
			table.Integer("log_type")
			table.Text("log_desc")
			table.Text("address")
			table.Text("agent")
			table.Text("device")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519042624CreateLogTable) Down() error {
 	return facades.Schema().DropIfExists("log")
}

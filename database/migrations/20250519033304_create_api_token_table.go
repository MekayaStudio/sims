package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519033304CreateApiTokenTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519033304CreateApiTokenTable) Signature() string {
	return "20250519033304_create_api_token_table"
}

// Up Run the migrations.
func (r *M20250519033304CreateApiTokenTable) Up() error {
	if !facades.Schema().HasTable("api_token") {
		return facades.Schema().Create("api_token", func(table schema.Blueprint) {
			table.ID("id_api")
			table.DateTime("timestamp")
			table.Integer("id_user")
			table.Text("address")
			table.Text("agent")
			table.Text("device")
			table.Text("token")
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519033304CreateApiTokenTable) Down() error {
 	return facades.Schema().DropIfExists("api_token")
}

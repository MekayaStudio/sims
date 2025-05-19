package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519041158CreateCbtTokenTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519041158CreateCbtTokenTable) Signature() string {
	return "20250519041158_create_cbt_token_table"
}

// Up Run the migrations.
func (r *M20250519041158CreateCbtTokenTable) Up() error {
	if !facades.Schema().HasTable("cbt_token") {
		return facades.Schema().Create("cbt_token", func(table schema.Blueprint) {
			table.Increments("id_token")
			table.String("token", 6)
			table.Integer("auto")
			table.Integer("jarak").Default(0)
			table.String("updated", 20).Default("")
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519041158CreateCbtTokenTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_token")
}

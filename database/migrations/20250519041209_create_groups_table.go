package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519041209CreateGroupsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519041209CreateGroupsTable) Signature() string {
	return "20250519041209_create_groups_table"
}

// Up Run the migrations.
func (r *M20250519041209CreateGroupsTable) Up() error {
	if !facades.Schema().HasTable("groups") {
		return facades.Schema().Create("groups", func(table schema.Blueprint) {
			table.MediumIncrements("id")
			table.String("name", 20)
			table.String("description", 100)
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519041209CreateGroupsTable) Down() error {
 	return facades.Schema().DropIfExists("groups")
}

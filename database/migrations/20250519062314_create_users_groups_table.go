package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519062314CreateUsersGroupsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519062314CreateUsersGroupsTable) Signature() string {
	return "20250519062314_create_users_groups_table"
}

// Up Run the migrations.
func (r *M20250519062314CreateUsersGroupsTable) Up() error {
	if !facades.Schema().HasTable("users_groups") {
		return facades.Schema().Create("users_groups", func(table schema.Blueprint) {
			table.ID("id")
			table.UnsignedInteger("user_id")
			table.UnsignedMediumInteger("group_id")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519062314CreateUsersGroupsTable) Down() error {
 	return facades.Schema().DropIfExists("users_groups")
}

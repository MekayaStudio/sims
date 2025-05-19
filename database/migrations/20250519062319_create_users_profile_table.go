package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519062319CreateUsersProfileTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519062319CreateUsersProfileTable) Signature() string {
	return "20250519062319_create_users_profile_table"
}

// Up Run the migrations.
func (r *M20250519062319CreateUsersProfileTable) Up() error {
	if !facades.Schema().HasTable("users_profile") {
		return facades.Schema().Create("users_profile", func(table schema.Blueprint) {
			table.Increments("id_user")
			table.Text("nama_lengkap")
			table.Text("jabatan").Nullable()
			table.Integer("level_access").Default(0)
			table.Text("foto").Nullable()
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519062319CreateUsersProfileTable) Down() error {
 	return facades.Schema().DropIfExists("users_profile")
}

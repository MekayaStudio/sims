package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20240915060148CreateUsersTable struct {
}

// Signature The unique signature for the migration.
func (r *M20240915060148CreateUsersTable) Signature() string {
	return "20240915060148_create_users_table"
}

// Up Run the migrations.
func (r *M20240915060148CreateUsersTable) Up() error {
	return facades.Schema().Create("users", func(table schema.Blueprint) {
		table.ID("id")
		table.String("ip_address", 45)
		table.String("username", 100).Nullable()
		table.String("password", 255)
		table.String("email", 254).Nullable()
		table.String("activation_selector", 255).Nullable()
		table.String("activation_code", 255).Nullable()
		table.String("forgotten_password_selector", 255).Nullable()
		table.String("forgotten_password_code", 255).Nullable()
		table.UnsignedInteger("forgotten_password_time").Nullable()
		table.String("remember_selector", 255).Nullable()
		table.String("remember_code", 255).Nullable()
		table.UnsignedInteger("created_on")
		table.UnsignedInteger("last_login").Nullable()
		table.UnsignedTinyInteger("active").Nullable()
		table.String("first_name", 50).Nullable()
		table.String("last_name", 50).Nullable()
		table.String("company", 100).Nullable()
		table.String("phone", 20).Nullable()
		table.Unique("username")
	})
}

// Down Reverse the migrations.
func (r *M20240915060148CreateUsersTable) Down() error {
	return facades.Schema().DropIfExists("users")
}

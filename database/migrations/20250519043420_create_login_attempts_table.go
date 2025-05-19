package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519043420CreateLoginAttemptsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519043420CreateLoginAttemptsTable) Signature() string {
	return "20250519043420_create_login_attempts_table"
}

// Up Run the migrations.
func (r *M20250519043420CreateLoginAttemptsTable) Up() error {
	if !facades.Schema().HasTable("login_attempts") {
		return facades.Schema().Create("login_attempts", func(table schema.Blueprint) {
			table.Increments("id")
			table.String("ip_address", 45)
			table.String("login", 100)
			table.Integer("time").Unsigned()
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519043420CreateLoginAttemptsTable) Down() error {
 	return facades.Schema().DropIfExists("login_attempts")
}

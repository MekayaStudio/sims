package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519045120CreateRunningTextTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519045120CreateRunningTextTable) Signature() string {
	return "20250519045120_create_running_text_table"
}

// Up Run the migrations.
func (r *M20250519045120CreateRunningTextTable) Up() error {
	if !facades.Schema().HasTable("running_text") {
		return facades.Schema().Create("running_text", func(table schema.Blueprint) {
			table.Increments("id_text")
			table.String("text", 255).Nullable()
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519045120CreateRunningTextTable) Down() error {
 	return facades.Schema().DropIfExists("running_text")
}

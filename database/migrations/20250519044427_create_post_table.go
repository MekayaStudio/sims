package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519044427CreatePostTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519044427CreatePostTable) Signature() string {
	return "20250519044427_create_post_table"
}

// Up Run the migrations.
func (r *M20250519044427CreatePostTable) Up() error {
	if !facades.Schema().HasTable("post") {
		return facades.Schema().Create("post", func(table schema.Blueprint) {
			table.Increments("id_post")
			table.Integer("dari").Nullable()
			table.Integer("dari_group").Nullable()
			table.String("kepada", 50)
			table.LongText("text")
			table.DateTime("tanggal")
			table.DateTime("updated")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519044427CreatePostTable) Down() error {
 	return facades.Schema().DropIfExists("post")
}

package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519044432CreatePostCommentsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519044432CreatePostCommentsTable) Signature() string {
	return "20250519044432_create_post_comments_table"
}

// Up Run the migrations.
func (r *M20250519044432CreatePostCommentsTable) Up() error {
	if !facades.Schema().HasTable("post_comments") {
		return facades.Schema().Create("post_comments", func(table schema.Blueprint) {
			table.Increments("id_comment")
			table.Integer("id_post").Nullable()
			table.Integer("dari").Nullable()
			table.Integer("dari_group").Nullable()
			table.LongText("text")
			table.DateTime("tanggal")
			table.DateTime("updated")
			table.String("type", 1)
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519044432CreatePostCommentsTable) Down() error {
 	return facades.Schema().DropIfExists("post_comments")
}

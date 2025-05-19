package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519044438CreatePostReplyTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519044438CreatePostReplyTable) Signature() string {
	return "20250519044438_create_post_reply_table"
}

// Up Run the migrations.
func (r *M20250519044438CreatePostReplyTable) Up() error {
	if !facades.Schema().HasTable("post_reply") {
		return facades.Schema().Create("post_reply", func(table schema.Blueprint) {
			table.Increments("id_reply")
			table.Integer("id_comment").Nullable()
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
func (r *M20250519044438CreatePostReplyTable) Down() error {
 	return facades.Schema().DropIfExists("post_reply")
}

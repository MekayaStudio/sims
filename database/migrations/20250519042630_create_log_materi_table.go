package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519042630CreateLogMateriTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519042630CreateLogMateriTable) Signature() string {
	return "20250519042630_create_log_materi_table"
}

// Up Run the migrations.
func (r *M20250519042630CreateLogMateriTable) Up() error {
	if !facades.Schema().HasTable("log_materi") {
		return facades.Schema().Create("log_materi", func(table schema.Blueprint) {
			table.String("id_log", 50)
			table.DateTime("log_time")
			table.Integer("id_siswa").Nullable()
			table.Integer("jam_ke")
			table.String("id_materi", 50)
			table.Integer("id_mapel").Nullable()
			table.Integer("log_type")
			table.Text("log_desc")
			table.LongText("text").Nullable()
			table.MediumText("file").Nullable()
			table.String("nilai", 3).Nullable()
			table.MediumText("catatan").Nullable()
			table.Text("address")
			table.Text("agent")
			table.Text("device")
			table.String("finish_time", 20).Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519042630CreateLogMateriTable) Down() error {
 	return facades.Schema().DropIfExists("log_materi")
}

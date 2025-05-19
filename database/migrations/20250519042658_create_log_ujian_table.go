package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519042658CreateLogUjianTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519042658CreateLogUjianTable) Signature() string {
	return "20250519042658_create_log_ujian_table"
}

// Up Run the migrations.
func (r *M20250519042658CreateLogUjianTable) Up() error {
	if !facades.Schema().HasTable("log_ujian") {
		return facades.Schema().Create("log_ujian", func(table schema.Blueprint) {
			table.Increments("id_log")
			table.DateTime("log_time")
			table.Integer("id_siswa").Nullable()
			table.Integer("id_jadwal").Nullable()
			table.Integer("log_type")
			table.Text("log_desc")
			table.Text("address")
			table.Text("agent")
			table.Text("device")
			table.Integer("reset")
			table.String("finish_time", 20).Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519042658CreateLogUjianTable) Down() error {
 	return facades.Schema().DropIfExists("log_ujian")
}

package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519034518CreateCbtSesiTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519034518CreateCbtSesiTable) Signature() string {
	return "20250519034518_create_cbt_sesi_table"
}

// Up Run the migrations.
func (r *M20250519034518CreateCbtSesiTable) Up() error {
	if !facades.Schema().HasTable("cbt_sesi") {
		return facades.Schema().Create("cbt_sesi", func(table schema.Blueprint) {
			table.Increments("id_sesi")
			table.String("nama_sesi", 50)
			table.String("kode_sesi", 10)
			table.Time("waktu_mulai")
			table.Time("waktu_akhir")
			table.Integer("aktif")
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519034518CreateCbtSesiTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_sesi")
}

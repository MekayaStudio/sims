package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519034442CreateCbtPengawasTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519034442CreateCbtPengawasTable) Signature() string {
	return "20250519034442_create_cbt_pengawas_table"
}

// Up Run the migrations.
func (r *M20250519034442CreateCbtPengawasTable) Up() error {
	if !facades.Schema().HasTable("cbt_pengawas") {
		return facades.Schema().Create("cbt_pengawas", func(table schema.Blueprint) {
			table.String("id_pengawas", 50)
			table.String("id_jadwal", 50)
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.String("id_ruang", 50)
			table.String("id_sesi", 50)
			table.String("id_guru", 50)
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519034442CreateCbtPengawasTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_pengawas")
}

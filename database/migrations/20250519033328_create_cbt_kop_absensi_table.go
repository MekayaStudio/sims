package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519033328CreateCbtKopAbsensiTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519033328CreateCbtKopAbsensiTable) Signature() string {
	return "20250519033328_create_cbt_kop_absensi_table"
}

// Up Run the migrations.
func (r *M20250519033328CreateCbtKopAbsensiTable) Up() error {
	if !facades.Schema().HasTable("cbt_kop_absensi") {
		return facades.Schema().Create("cbt_kop_absensi", func(table schema.Blueprint) {
			table.ID("id_kop")
			table.String("header_1", 100).Nullable()
			table.String("header_2", 100).Nullable()
			table.String("header_3", 100).Nullable()
			table.String("header_4", 100).Nullable()
			table.String("proktor", 100).Nullable()
			table.String("pengawas_1", 100).Nullable()
			table.String("pengawas_2", 100).Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519033328CreateCbtKopAbsensiTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_kop_absensi")
}

package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519033325CreateCbtKelasRuangTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519033325CreateCbtKelasRuangTable) Signature() string {
	return "20250519033325_create_cbt_kelas_ruang_table"
}

// Up Run the migrations.
func (r *M20250519033325CreateCbtKelasRuangTable) Up() error {
	if !facades.Schema().HasTable("cbt_kelas_ruang") {
		return facades.Schema().Create("cbt_kelas_ruang", func(table schema.Blueprint) {
			table.String("id_kelas_ruang", 50)
			table.Integer("id_kelas").Nullable()
			table.Integer("id_ruang").Nullable()
			table.Integer("id_sesi").Default(0).Nullable()
			table.Integer("id_tp").Nullable()
			table.Integer("id_smt").Nullable()
			table.Integer("set_siswa").Default(0).Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519033325CreateCbtKelasRuangTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_kelas_ruang")
}

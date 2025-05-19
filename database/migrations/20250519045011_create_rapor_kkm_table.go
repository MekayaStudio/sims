package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519045011CreateRaporKkmTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519045011CreateRaporKkmTable) Signature() string {
	return "20250519045011_create_rapor_kkm_table"
}

// Up Run the migrations.
func (r *M20250519045011CreateRaporKkmTable) Up() error {
	if !facades.Schema().HasTable("rapor_kkm") {
		return facades.Schema().Create("rapor_kkm", func(table schema.Blueprint) {
			table.Increments("id_kkm")
			table.Integer("kkm").Default(0)
			table.Integer("bobot_ph").Default(0)
			table.Integer("bobot_pts").Default(0)
			table.Integer("bobot_pas").Default(0)
			table.Integer("bobot_absen").Default(0)
			table.Integer("beban_jam").Default(0)
			table.Integer("id_tp").Default(0)
			table.Integer("id_smt").Default(0)
			table.Integer("jenis")
			table.Integer("id_kelas").Nullable()
			table.Integer("id_mapel").Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519045011CreateRaporKkmTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_kkm")
}

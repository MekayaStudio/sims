package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519044508CreateRaporDataSikapTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519044508CreateRaporDataSikapTable) Signature() string {
	return "20250519044508_create_rapor_data_sikap_table"
}

// Up Run the migrations.
func (r *M20250519044508CreateRaporDataSikapTable) Up() error {
	if !facades.Schema().HasTable("rapor_data_sikap") {
		return facades.Schema().Create("rapor_data_sikap", func(table schema.Blueprint) {
			table.Increments("id_sikap")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("id_kelas").Nullable()
			table.Integer("jenis")
			table.Integer("kode")
			table.String("sikap", 100)
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519044508CreateRaporDataSikapTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_data_sikap")
}

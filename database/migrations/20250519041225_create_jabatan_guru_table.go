package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519041225CreateJabatanGuruTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519041225CreateJabatanGuruTable) Signature() string {
	return "20250519041225_create_jabatan_guru_table"
}

// Up Run the migrations.
func (r *M20250519041225CreateJabatanGuruTable) Up() error {
	if !facades.Schema().HasTable("jabatan_guru") {
		return facades.Schema().Create("jabatan_guru", func(table schema.Blueprint) {
			table.String("id_jabatan_guru", 50)
			table.Integer("id_guru").Nullable()
			table.Integer("id_jabatan")
			table.Integer("id_kelas").Default(0)
			table.LongText("mapel_kelas").Nullable()
			table.LongText("ekstra_kelas").Nullable()
			table.Integer("id_tp")
			table.Integer("id_smt")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519041225CreateJabatanGuruTable) Down() error {
 	return facades.Schema().DropIfExists("jabatan_guru")
}

package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519041248CreateKelasEkstraTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519041248CreateKelasEkstraTable) Signature() string {
	return "20250519041248_create_kelas_ekstra_table"
}

// Up Run the migrations.
func (r *M20250519041248CreateKelasEkstraTable) Up() error {
	if !facades.Schema().HasTable("kelas_ekstra") {
		return facades.Schema().Create("kelas_ekstra", func(table schema.Blueprint) {
			table.String("id_kelas_ekstra", 50)
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("id_kelas").Nullable()
			table.LongText("ekstra")
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519041248CreateKelasEkstraTable) Down() error {
 	return facades.Schema().DropIfExists("kelas_ekstra")
}

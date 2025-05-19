package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519045113CreateRaporPrestasiTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519045113CreateRaporPrestasiTable) Signature() string {
	return "20250519045113_create_rapor_prestasi_table"
}

// Up Run the migrations.
func (r *M20250519045113CreateRaporPrestasiTable) Up() error {
	if !facades.Schema().HasTable("rapor_prestasi") {
		return facades.Schema().Create("rapor_prestasi", func(table schema.Blueprint) {
			table.Increments("id_ranking")
			table.Integer("id_kelas").Nullable()
			table.Integer("id_siswa").Nullable()
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("ranking")
			table.String("deskripsi", 100)
			table.String("p1", 100)
			table.String("p1_desk", 100)
			table.String("p2", 100)
			table.String("p2_desk", 100)
			table.String("p3", 100)
			table.String("p3_desk", 100)
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519045113CreateRaporPrestasiTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_prestasi")
}

package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519034422CreateCbtNilaiTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519034422CreateCbtNilaiTable) Signature() string {
	return "20250519034422_create_cbt_nilai_table"
}

// Up Run the migrations.
func (r *M20250519034422CreateCbtNilaiTable) Up() error {
	if !facades.Schema().HasTable("cbt_nilai") {
		return facades.Schema().Create("cbt_nilai", func(table schema.Blueprint) {
			table.ID("id_nilai")
			table.Integer("pg_benar").Default(0)
			table.String("pg_nilai", 10).Default("0")
			table.String("essai_nilai", 10).Default("0")
			table.Integer("id_siswa").Nullable()
			table.Integer("id_jadwal").Nullable()
			table.String("kompleks_nilai", 10).Default("0")
			table.String("jodohkan_nilai", 10).Default("0")
			table.String("isian_nilai", 10).Default("0")
			table.Integer("dikoreksi").Default(0)
			table.DateTime("time_create").Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519034422CreateCbtNilaiTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_nilai")
}

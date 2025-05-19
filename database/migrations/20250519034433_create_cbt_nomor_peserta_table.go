package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519034433CreateCbtNomorPesertaTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519034433CreateCbtNomorPesertaTable) Signature() string {
	return "20250519034433_create_cbt_nomor_peserta_table"
}

// Up Run the migrations.
func (r *M20250519034433CreateCbtNomorPesertaTable) Up() error {
	if !facades.Schema().HasTable("cbt_nomor_peserta") {
		return facades.Schema().Create("cbt_nomor_peserta", func(table schema.Blueprint) {
			table.ID("id_nomor")
			table.Integer("id_siswa").Nullable()
			table.Integer("id_tp")
			table.Integer("id_smt").Default(1)
			table.String("nomor_peserta", 20)
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250519034433CreateCbtNomorPesertaTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_nomor_peserta")
}

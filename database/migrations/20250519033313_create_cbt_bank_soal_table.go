package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519033313CreateCbtBankSoalTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519033313CreateCbtBankSoalTable) Signature() string {
	return "20250519033313_create_cbt_bank_soal_table"
}

// Up Run the migrations.
func (r *M20250519033313CreateCbtBankSoalTable) Up() error {
	if !facades.Schema().HasTable("cbt_bank_soal") {
		return facades.Schema().Create("cbt_bank_soal", func(table schema.Blueprint) {
			table.ID("id_bank")
			table.Integer("bank_jenis_id").Default(0)
			table.String("bank_kode", 255).Default("0")
			table.String("bank_level", 225)
			table.LongText("bank_kelas")
			table.Integer("bank_mapel_id")
			table.Integer("bank_jurusan_id").Default(0)
			table.Integer("bank_guru_id")
			table.String("bank_nama", 250)
			table.Integer("kkm").Default(0)
			table.Integer("jml_soal").Default(0)
			table.Integer("jml_esai").Default(0)
			table.Integer("tampil_pg").Default(0)
			table.Integer("tampil_esai").Default(0)
			table.Integer("bobot_pg").Default(0)
			table.Integer("bobot_esai").Default(0)
			table.Integer("opsi").Default(0)
			table.Timestamp("date")
			table.Integer("status").Default(0)
			table.String("soal_agama", 20)
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.LongText("deskripsi")
			table.Integer("jml_kompleks").Default(0)
			table.Integer("tampil_kompleks").Default(0)
			table.Integer("bobot_kompleks").Default(0)
			table.Integer("jml_jodohkan").Default(0)
			table.Integer("tampil_jodohkan").Default(0)
			table.Integer("bobot_jodohkan").Default(0)
			table.Integer("jml_isian").Default(0)
			table.Integer("tampil_isian").Default(0)
			table.Integer("bobot_isian").Default(0)
			table.Integer("status_soal").Default(0)
			table.Unique("bank_kode")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519033313CreateCbtBankSoalTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_bank_soal")
}

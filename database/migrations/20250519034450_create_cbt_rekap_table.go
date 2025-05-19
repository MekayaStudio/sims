package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519034450CreateCbtRekapTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519034450CreateCbtRekapTable) Signature() string {
	return "20250519034450_create_cbt_rekap_table"
}

// Up Run the migrations.
func (r *M20250519034450CreateCbtRekapTable) Up() error {
	if !facades.Schema().HasTable("cbt_rekap") {
		return facades.Schema().Create("cbt_rekap", func(table schema.Blueprint) {
			table.Increments("id_rekap")
			table.Integer("id_tp")
			table.String("tp", 20)
			table.Integer("id_smt")
			table.String("smt", 20)
			table.String("id_jadwal", 250)
			table.String("id_jenis", 250)
			table.String("kode_jenis", 20)
			table.String("id_bank", 250)
			table.MediumText("bank_kelas")
			table.String("bank_kode", 20)
			table.Integer("bank_level")
			table.String("id_mapel", 250)
			table.String("nama_mapel", 100)
			table.String("kode", 50)
			table.String("tgl_mulai", 22)
			table.String("tgl_selesai", 22)
			table.Integer("tampil_pg")
			table.LongText("jawaban_pg")
			table.Integer("tampil_esai")
			table.LongText("jawaban_esai")
			table.Integer("bobot_pg")
			table.Integer("bobot_esai")
			table.String("id_guru", 250)
			table.String("nama_guru", 100)
			table.MediumText("nama_kelas").Nullable()
			table.LongText("soal_kompleks").Nullable()
			table.LongText("soal_jodohkan").Nullable()
			table.LongText("soal_isian").Nullable()
			table.LongText("soal_essai").Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519034450CreateCbtRekapTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_rekap")
}

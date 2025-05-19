package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519034500CreateCbtRekapNilaiTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519034500CreateCbtRekapNilaiTable) Signature() string {
	return "20250519034500_create_cbt_rekap_nilai_table"
}

// Up Run the migrations.
func (r *M20250519034500CreateCbtRekapNilaiTable) Up() error {
	if !facades.Schema().HasTable("cbt_rekap_nilai") {
		return facades.Schema().Create("cbt_rekap_nilai", func(table schema.Blueprint) {
			table.Increments("id_rekap_nilai")
			table.Integer("id_jadwal").Nullable()
			table.Integer("id_tp")
			table.String("tp", 20)
			table.Integer("id_smt")
			table.String("smt", 20)
			table.Integer("id_jenis")
			table.String("kode_jenis", 20)
			table.Integer("id_bank").Nullable()
			table.Integer("id_mapel").Nullable()
			table.Integer("id_siswa").Nullable()
			table.Integer("id_kelas").Default(0)
			table.String("kelas", 20)
			table.String("mulai", 20)
			table.String("selesai", 20)
			table.String("durasi", 20)
			table.Integer("bobot_pg")
			table.LongText("jawaban_pg")
			table.String("nilai_pg", 10)
			table.Integer("bobot_esai")
			table.LongText("jawaban_esai")
			table.String("nilai_esai", 10)
			table.Integer("id_guru").Nullable()
			table.String("nama_siswa", 100).Nullable()
			table.String("no_peserta", 100).Nullable()
			table.LongText("soal_kompleks").Nullable()
			table.LongText("soal_jodohkan").Nullable()
			table.LongText("soal_isian").Nullable()
			table.LongText("soal_essai").Nullable()
			table.DateTime("time_create").Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519034500CreateCbtRekapNilaiTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_rekap_nilai")
}

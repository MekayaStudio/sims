package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519041150CreateCbtSoalSiswaTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519041150CreateCbtSoalSiswaTable) Signature() string {
	return "20250519041150_create_cbt_soal_siswa_table"
}

// Up Run the migrations.
func (r *M20250519041150CreateCbtSoalSiswaTable) Up() error {
	if !facades.Schema().HasTable("cbt_soal_siswa") {
		return facades.Schema().Create("cbt_soal_siswa", func(table schema.Blueprint) {
			table.String("id_soal_siswa", 15)
			table.Integer("id_bank").Nullable()
			table.Integer("id_jadwal").Nullable()
			table.Integer("id_soal").Nullable()
			table.Integer("id_siswa").Nullable()
			table.Integer("jenis_soal")
			table.Integer("no_soal_alias")
			table.String("opsi_alias_a", 1).Nullable()
			table.String("opsi_alias_b", 1).Nullable()
			table.String("opsi_alias_c", 1).Nullable()
			table.String("opsi_alias_d", 1).Nullable()
			table.String("opsi_alias_e", 1).Nullable()
			table.LongText("jawaban_alias").Nullable()
			table.LongText("jawaban_siswa").Nullable()
			table.LongText("jawaban_benar").Nullable()
			table.Integer("point_essai").Default(0)
			table.Integer("soal_end").Default(0)
			table.String("point_soal", 5).Default("0")
			table.String("nilai_koreksi", 5).Default("0")
			table.Integer("nilai_otomatis").Default(0).Comment("0=otomatis, 1=dari guru")
			table.DateTime("time_create").Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519041150CreateCbtSoalSiswaTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_soal_siswa")
}

package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519041141CreateCbtSoalTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519041141CreateCbtSoalTable) Signature() string {
	return "20250519041141_create_cbt_soal_table"
}

// Up Run the migrations.
func (r *M20250519041141CreateCbtSoalTable) Up() error {
	if !facades.Schema().HasTable("cbt_soal") {
		return facades.Schema().Create("cbt_soal", func(table schema.Blueprint) {
			table.Increments("id_soal")
			table.Integer("bank_id").Nullable()
			table.Integer("mapel_id").Default(0)
			table.Integer("jenis").Comment("1=ganda, 2=ganda kompleks, 3=menjodohkan, 4=isian singkat, 5=uraian")
			table.Integer("nomor_soal").Default(0)
			table.String("file", 255).Nullable()
			table.MediumText("file1").Nullable()
			table.String("tipe_file", 50).Nullable()
			table.LongText("soal").Nullable()
			table.LongText("opsi_a").Nullable()
			table.LongText("opsi_b").Nullable()
			table.LongText("opsi_c").Nullable()
			table.LongText("opsi_d").Nullable()
			table.LongText("opsi_e").Nullable()
			table.String("file_a", 255).Nullable()
			table.String("file_b", 255).Nullable()
			table.String("file_c", 255).Nullable()
			table.String("file_d", 255).Nullable()
			table.String("file_e", 255).Nullable()
			table.LongText("jawaban").Nullable()
			table.Integer("created_on").Nullable()
			table.Integer("updated_on").Nullable()
			table.Integer("tampilkan").Default(0)
			table.LongText("deskripsi")
			table.Integer("kesulitan").Default(1).Comment("tingkat kesulitan 1-10")
			table.Integer("timer").Default(0).Comment("0=tidak, 1=ya")
			table.Integer("timer_menit").Default(0)
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519041141CreateCbtSoalTable) Down() error {
 	return facades.Schema().DropIfExists("cbt_soal")
}

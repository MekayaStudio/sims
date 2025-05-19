package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519033307CreateBukuIndukTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519033307CreateBukuIndukTable) Signature() string {
	return "20250519033307_create_buku_induk_table"
}

// Up Run the migrations.
func (r *M20250519033307CreateBukuIndukTable) Up() error {
	if !facades.Schema().HasTable("buku_induk") {
		return facades.Schema().Create("buku_induk", func(table schema.Blueprint) {
			table.ID("id_siswa")
			table.String("uid", 50)
			table.String("rombel_awal", 50)
			table.String("nama_panggilan", 50)
			table.String("bahasa", 50)
			table.Integer("jml_saudara_kandung").Default(0)
			table.Integer("jml_saudara_tiri").Default(0)
			table.Integer("jml_saudara_angkat").Default(0)
			table.Integer("yatim").Default(0)
			table.String("tinggal_bersama", 1).Default("1")
			table.String("jarak", 10)
			table.String("gol_darah", 4)
			table.MediumText("penyakit")
			table.String("kelainan_fisik", 100)
			table.LongText("kegemaran")
			table.LongText("beasiswa")
			table.String("no_ijazah_sebelumnya", 50)
			table.String("tahun_lulus_sebelumnya", 10)
			table.String("pindahan_dari", 100)
			table.String("alasan_kepindahan", 200)
			table.String("agama_ayah", 20)
			table.String("tempat_lahir_ayah", 50)
			table.String("wn_ayah", 50)
			table.String("penghasilan_ayah", 50)
			table.String("hidup_meninggal_ayah", 50)
			table.String("agama_ibu", 50)
			table.String("tempat_lahir_ibu", 50)
			table.String("wn_ibu", 50)
			table.String("penghasilan_ibu", 50)
			table.String("hidup_meninggal_ibu", 50)
			table.String("tempat_lahir_wali", 50)
			table.String("agama_wali", 20)
			table.String("wn_wali", 50)
			table.String("penghasilan_wali", 10)
			table.Integer("status").Default(1)
			table.Integer("tahun_lulus")
			table.String("no_ijazah", 50)
			table.String("kelas_akhir", 50)
			table.String("lanjut_ke", 50)
			table.String("pindah_ke", 100)
			table.String("alasan_pindah", 100)
			table.String("tgl_pindah", 20)
			table.String("bekerja_di", 100)
			table.LongText("catatan_penting")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519033307CreateBukuIndukTable) Down() error {
 	return facades.Schema().DropIfExists("buku_induk")
}

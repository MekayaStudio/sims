package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519043502CreateMasterSiswaTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519043502CreateMasterSiswaTable) Signature() string {
	return "20250519043502_create_master_siswa_table"
}

// Up Run the migrations.
func (r *M20250519043502CreateMasterSiswaTable) Up() error {
	if !facades.Schema().HasTable("master_siswa") {
		return facades.Schema().Create("master_siswa", func(table schema.Blueprint) {
			table.Increments("id_siswa")
			table.BigInteger("nisn").Unsigned()
			table.String("nis", 20)
			table.String("nama", 50)
			table.Char("jenis_kelamin", 1)
			table.String("username", 50)
			table.Text("password")
			table.Integer("kelas_awal")
			table.String("tahun_masuk", 30)
			table.String("sekolah_asal", 100)
			table.String("tempat_lahir", 100)
			table.String("tanggal_lahir", 30)
			table.String("agama", 10)
			table.String("hp", 15)
			table.String("email", 254)
			table.String("foto", 255).Default("siswa.png")
			table.Integer("anak_ke")
			table.Char("status_keluarga", 1)
			table.MediumText("alamat")
			table.String("rt", 5)
			table.String("rw", 5)
			table.String("kelurahan", 100)
			table.String("kecamatan", 100)
			table.String("kabupaten", 100)
			table.String("provinsi", 100)
			table.Integer("kode_pos")
			table.String("nama_ayah", 150)
			table.String("tgl_lahir_ayah", 50)
			table.String("pendidikan_ayah", 50)
			table.String("pekerjaan_ayah", 100)
			table.String("nohp_ayah", 20)
			table.LongText("alamat_ayah")
			table.String("nama_ibu", 50)
			table.String("tgl_lahir_ibu", 50)
			table.String("pendidikan_ibu", 50)
			table.String("pekerjaan_ibu", 100)
			table.String("nohp_ibu", 20)
			table.LongText("alamat_ibu")
			table.String("nama_wali", 150)
			table.String("tgl_lahir_wali", 50)
			table.String("pendidikan_wali", 50)
			table.String("pekerjaan_wali", 100)
			table.String("nohp_wali", 20)
			table.LongText("alamat_wali")
			table.String("nik", 30)
			table.String("warga_negara", 20)
			table.String("uid", 255)
			// Primary and unique keys will be handled in a later step if needed
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519043502CreateMasterSiswaTable) Down() error {
 	return facades.Schema().DropIfExists("master_siswa")
}

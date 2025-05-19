package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519043432CreateMasterGuruTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519043432CreateMasterGuruTable) Signature() string {
	return "20250519043432_create_master_guru_table"
}

// Up Run the migrations.
func (r *M20250519043432CreateMasterGuruTable) Up() error {
	if !facades.Schema().HasTable("master_guru") {
		return facades.Schema().Create("master_guru", func(table schema.Blueprint) {
			table.Increments("id_guru")
			table.Integer("id_user")
			table.Char("nip", 30)
			table.String("nama_guru", 50)
			table.String("email", 254)
			table.String("kode_guru", 6)
			table.String("username", 50)
			table.MediumText("password")
			table.String("no_ktp", 16)
			table.String("tempat_lahir", 30)
			table.Date("tgl_lahir")
			table.String("jenis_kelamin", 10)
			table.String("agama", 10)
			table.String("no_hp", 13)
			table.String("alamat_jalan", 255)
			table.String("rt_rw", 8)
			table.String("dusun", 50)
			table.String("kelurahan", 50)
			table.String("kecamatan", 30)
			table.String("kabupaten", 30)
			table.String("provinsi", 30)
			table.Integer("kode_pos")
			table.String("kewarganegaraan", 10)
			table.String("nuptk", 20)
			table.String("jenis_ptk", 50)
			table.String("tgs_tambahan", 50)
			table.String("status_pegawai", 50)
			table.String("status_aktif", 20)
			table.String("status_nikah", 20)
			table.Date("tmt")
			table.String("keahlian_isyarat", 10)
			table.String("npwp", 16)
			table.LongText("foto")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519043432CreateMasterGuruTable) Down() error {
 	return facades.Schema().DropIfExists("master_guru")
}

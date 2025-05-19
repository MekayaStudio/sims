package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519062303CreateSettingTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519062303CreateSettingTable) Signature() string {
	return "20250519062303_create_setting_table"
}

// Up Run the migrations.
func (r *M20250519062303CreateSettingTable) Up() error {
	if !facades.Schema().HasTable("setting") {
		return facades.Schema().Create("setting", func(table schema.Blueprint) {
			table.Increments("id_setting")
			table.String("kode_sekolah", 10).Nullable()
			table.String("sekolah", 50).Nullable()
			table.String("npsn", 10).Nullable()
			table.String("nss", 20).Nullable()
			table.Integer("jenjang").Nullable()
			table.String("kepsek", 50).Nullable()
			table.String("nip", 30).Nullable()
			table.MediumText("tanda_tangan").Nullable()
			table.MediumText("alamat").Nullable()
			table.String("desa", 100).Nullable()
			table.String("kecamatan", 50).Nullable()
			table.String("kota", 30).Nullable()
			table.String("provinsi", 100).Nullable()
			table.Integer("kode_pos").Nullable()
			table.String("telp", 20).Nullable()
			table.String("fax", 20).Nullable()
			table.String("web", 50).Nullable()
			table.String("email", 50).Nullable()
			table.String("nama_aplikasi", 100).Nullable()
			table.MediumText("logo_kanan").Nullable()
			table.MediumText("logo_kiri").Nullable()
			table.String("versi", 10).Nullable()
			table.String("ip_server", 100).Nullable()
			table.String("waktu", 50).Nullable()
			table.String("server", 50).Nullable()
			table.String("id_server", 50).Nullable()
			table.String("sekolah_id", 50).Nullable()
			table.String("db_versi", 10).Nullable()
			table.String("satuan_pendidikan", 1).Nullable()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519062303CreateSettingTable) Down() error {
 	return facades.Schema().DropIfExists("setting")
}

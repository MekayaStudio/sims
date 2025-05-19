package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250519044446CreateRaporAdminSettingTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250519044446CreateRaporAdminSettingTable) Signature() string {
	return "20250519044446_create_rapor_admin_setting_table"
}

// Up Run the migrations.
func (r *M20250519044446CreateRaporAdminSettingTable) Up() error {
	if !facades.Schema().HasTable("rapor_admin_setting") {
		return facades.Schema().Create("rapor_admin_setting", func(table schema.Blueprint) {
			table.Increments("id_setting")
			table.Integer("id_tp")
			table.Integer("id_smt")
			table.Integer("kkm_tunggal")
			table.Integer("kkm").Nullable()
			table.Integer("bobot_ph").Nullable()
			table.Integer("bobot_pts").Nullable()
			table.Integer("bobot_pas").Nullable()
			table.Integer("bobot_absen").Nullable()
			table.String("tgl_rapor_akhir", 100).Nullable()
			table.String("tgl_rapor_kelas_akhir", 100).Nullable()
			table.String("tgl_rapor_pts", 100).Nullable()
			table.Integer("nip_kepsek")
			table.Integer("nip_walikelas")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250519044446CreateRaporAdminSettingTable) Down() error {
 	return facades.Schema().DropIfExists("rapor_admin_setting")
}

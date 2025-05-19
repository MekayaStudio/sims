package database

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/contracts/database/seeder"

	"goravel/database/migrations"
	"goravel/database/seeders"
)

type Kernel struct {
}

func (kernel Kernel) Migrations() []schema.Migration {
	return []schema.Migration{
		&migrations.M20250519062303CreateSettingTable{},
		&migrations.M20250519062314CreateUsersGroupsTable{},
		&migrations.M20250519062319CreateUsersProfileTable{},
		&migrations.M20240915060148CreateUsersTable{},
		&migrations.M20250519033301CreateApiSettingTable{},
		&migrations.M20250519033304CreateApiTokenTable{},
		&migrations.M20250519033307CreateBukuIndukTable{},
		&migrations.M20250519033310CreateBulanTable{},
		&migrations.M20250519033313CreateCbtBankSoalTable{},
		&migrations.M20250519033316CreateCbtDurasiSiswaTable{},
		&migrations.M20250519033319CreateCbtJadwalTable{},
		&migrations.M20250519033322CreateCbtJenisTable{},
		&migrations.M20250519033325CreateCbtKelasRuangTable{},
		&migrations.M20250519033328CreateCbtKopAbsensiTable{},
		&migrations.M20250519034405CreateCbtKopBeritaTable{},
		&migrations.M20250519034413CreateCbtKopKartuTable{},
		&migrations.M20250519034422CreateCbtNilaiTable{},
		&migrations.M20250519034433CreateCbtNomorPesertaTable{},
		&migrations.M20250519034442CreateCbtPengawasTable{},
		&migrations.M20250519034450CreateCbtRekapTable{},
		&migrations.M20250519034500CreateCbtRekapNilaiTable{},
		&migrations.M20250519034509CreateCbtRuangTable{},
		&migrations.M20250519034518CreateCbtSesiTable{},
		&migrations.M20250519034527CreateCbtSesiSiswaTable{},
		&migrations.M20250519041141CreateCbtSoalTable{},
		&migrations.M20250519041150CreateCbtSoalSiswaTable{},
		&migrations.M20250519041158CreateCbtTokenTable{},
		&migrations.M20250519041209CreateGroupsTable{},
		&migrations.M20250519041217CreateHariTable{},
		&migrations.M20250519041225CreateJabatanGuruTable{},
		&migrations.M20250519041231CreateKelasCatatanMapelTable{},
		&migrations.M20250519041240CreateKelasCatatanWaliTable{},
		&migrations.M20250519041248CreateKelasEkstraTable{},
		&migrations.M20250519041255CreateKelasJadwalKbmTable{},
		&migrations.M20250519042450CreateKelasJadwalMapelTable{},
		&migrations.M20250519042457CreateKelasJadwalMateriTable{},
		&migrations.M20250519042503CreateKelasMateriTable{},
		&migrations.M20250519042508CreateKelasSiswaTable{},
		&migrations.M20250519042517CreateKelasStrukturTable{},
		&migrations.M20250519042544CreateLevelGuruTable{},
		&migrations.M20250519042551CreateLevelKelasTable{},
		&migrations.M20250519042624CreateLogTable{},
		&migrations.M20250519042630CreateLogMateriTable{},
		&migrations.M20250519042658CreateLogUjianTable{},
		&migrations.M20250519043420CreateLoginAttemptsTable{},
		&migrations.M20250519043426CreateMasterEkstraTable{},
		&migrations.M20250519043432CreateMasterGuruTable{},
		&migrations.M20250519043437CreateMasterHariEfektifTable{},
		&migrations.M20250519043442CreateMasterJurusanTable{},
		&migrations.M20250519043447CreateMasterKelasTable{},
		&migrations.M20250519043452CreateMasterKelompokMapelTable{},
		&migrations.M20250519043457CreateMasterMapelTable{},
		&migrations.M20250519043502CreateMasterSiswaTable{},
		&migrations.M20250519043507CreateMasterSmtTable{},
		&migrations.M20250519044420CreateMasterTpTable{},
		&migrations.M20250519044427CreatePostTable{},
		&migrations.M20250519044432CreatePostCommentsTable{},
		&migrations.M20250519044438CreatePostReplyTable{},
		&migrations.M20250519044446CreateRaporAdminSettingTable{},
		&migrations.M20250519044452CreateRaporCatatanWaliTable{},
		&migrations.M20250519044457CreateRaporDataCatatanTable{},
		&migrations.M20250519044502CreateRaporDataFisikTable{},
		&migrations.M20250519044508CreateRaporDataSikapTable{},
		&migrations.M20250519044513CreateRaporFisikTable{},
		&migrations.M20250519045003CreateRaporKikdTable{},
		&migrations.M20250519045011CreateRaporKkmTable{},
		&migrations.M20250519045026CreateRaporNaikTable{},
		&migrations.M20250519045033CreateRaporNilaiAkhirTable{},
		&migrations.M20250519045041CreateRaporNilaiEkstraTable{},
		&migrations.M20250519045048CreateRaporNilaiHarianTable{},
		&migrations.M20250519045059CreateRaporNilaiPtsTable{},
		&migrations.M20250519045106CreateRaporNilaiSikapTable{},
		&migrations.M20250519045113CreateRaporPrestasiTable{},
		&migrations.M20250519045120CreateRunningTextTable{},
	}
}

func (kernel Kernel) Seeders() []seeder.Seeder {
	return []seeder.Seeder{
		&seeders.DatabaseSeeder{},
	}
}

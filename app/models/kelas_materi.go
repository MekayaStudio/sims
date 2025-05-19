package models

type KelasMateri struct {
	IDMateri    int    `gorm:"column:id_materi;primaryKey"`
	IDTp        int    `gorm:"column:id_tp"`
	IDSmt       int    `gorm:"column:id_smt"`
	KodeMateri  string `gorm:"column:kode_materi"`
	IDGuru      int    `gorm:"column:id_guru"`
	MateriKelas string `gorm:"column:materi_kelas"`
	IDMapel     int    `gorm:"column:id_mapel"`
	KodeMapel   string `gorm:"column:kode_mapel"`
	JudulMateri string `gorm:"column:judul_materi"`
	IsiMateri   string `gorm:"column:isi_materi"`
	File        string `gorm:"column:file"`
	LinkFile    string `gorm:"column:link_file"`
	TglMulai    string `gorm:"column:tgl_mulai"`
	CreatedOn   string `gorm:"column:created_on"`
	UpdatedOn   string `gorm:"column:updated_on"`
	Status      int    `gorm:"column:status"`
}

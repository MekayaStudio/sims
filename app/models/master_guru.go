package models

type MasterGuru struct {
	IDGuru      int    `gorm:"column:id_guru;primaryKey"`
	IDUser      int    `gorm:"column:id_user"`
	NIP         string `gorm:"column:nip"`
	NamaGuru    string `gorm:"column:nama_guru"`
	Email       string `gorm:"column:email"`
	KodeGuru    string `gorm:"column:kode_guru"`
	Username    string `gorm:"column:username"`
	Password    string `gorm:"column:password"`
	NoKTP       string `gorm:"column:no_ktp"`
	TempatLahir string `gorm:"column:tempat_lahir"`
	TglLahir    string `gorm:"column:tgl_lahir"`
	JenisKelamin string `gorm:"column:jenis_kelamin"`
	Agama       string `gorm:"column:agama"`
	NoHp        string `gorm:"column:no_hp"`
	AlamatJalan string `gorm:"column:alamat_jalan"`
	RtRw        string `gorm:"column:rt_rw"`
	Dusun       string `gorm:"column:dusun"`
	Kelurahan   string `gorm:"column:kelurahan"`
	Kecamatan   string `gorm:"column:kecamatan"`
	Kabupaten   string `gorm:"column:kabupaten"`
	Provinsi    string `gorm:"column:provinsi"`
	KodePos     int    `gorm:"column:kode_pos"`
}

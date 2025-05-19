package models

type KelasStruktur struct {
	IDKelas         int `gorm:"column:id_kelas;primaryKey"`
	Ketua           int `gorm:"column:ketua"`
	WakilKetua      int `gorm:"column:wakil_ketua"`
	Sekretaris1     int `gorm:"column:sekretaris_1"`
	Sekretaris2     int `gorm:"column:sekretaris_2"`
	Bendahara1      int `gorm:"column:bendahara_1"`
	Bendahara2      int `gorm:"column:bendahara_2"`
	SieEkstrakurikuler int `gorm:"column:sie_ekstrakurikuler"`
	SieUpacara      int `gorm:"column:sie_upacara"`
	SieOlahraga     int `gorm:"column:sie_olahraga"`
	SieKeagamaan    int `gorm:"column:sie_keagamaan"`
	SieKeamanan     int `gorm:"column:sie_keamanan"`
	SieKetertiban   int `gorm:"column:sie_ketertiban"`
	SieKebersihan   int `gorm:"column:sie_kebersihan"`
	SieKeindahan    int `gorm:"column:sie_keindahan"`
	SieKesehatan    int `gorm:"column:sie_kesehatan"`
	SieKekeluargaan int `gorm:"column:sie_kekeluargaan"`
	SieHumas        int `gorm:"column:sie_humas"`
}

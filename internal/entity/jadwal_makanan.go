package entity

type JadwalMakanan struct {
	ID       int    `json:"id" form:"id"`
	ID_Resep int    `json:"id_resep" form:"id_resep"`
	Tanggal  string `json:"tanggal" form:"tanggal"`
	Waktu    string `json:"waktu" form:"waktu"`
}

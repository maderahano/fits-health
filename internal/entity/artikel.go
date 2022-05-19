package entity

type Artikel struct {
	ID     int    `json:"id" form:"id"`
	Judul  string `json:"judul" form:"judul"`
	Jenis  string `json:"jenis" form:"jenis"`
	Gambar string `json:"gambar" form:"gambar"`
	Isi    string `json:"isi" form:"isi"`
}

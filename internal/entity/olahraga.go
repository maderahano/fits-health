package entity

type Olahraga struct {
	ID        int    `json:"id" form:"id"`
	Nama      string `json:"nama" form:"nama"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
}

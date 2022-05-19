package entity

type ResepMakanan struct {
	ID        int    `json:"id" form:"id"`
	Nama      string `json:"nama" form:"nama"`
	Durasi    int    `json:"durasi" form:"durasi"`
	Kalori    int    `json:"kalori" form:"kalori"`
	Porsi     int    `json:"porsi" form:"porsi"`
	Bahan     string `json:"bahan" form:"bahan"`
	Persiapan string `json:"persiapan" form:"persiapan"`
	Langkah   string `json:"langkah" form:"langkah"`
}

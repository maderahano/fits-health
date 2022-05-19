package entity

type ProfilKesehatan struct {
	ID            int    `json:"id" form:"id"`
	BeratBadan    int    `json:"berat_badan" form:"berat_badan"`
	TinggiBadan   int    `json:"tinggi_badan" form:"tinggi_badan"`
	GolonganDarah string `json:"golongan_darah" form:"golongan_darah"`
	Tensi         string `json:"tensi" form:"tensi"`
}

package entity

type User struct {
	ID           int    `json:"id" form:"id"`
	ID_IMT       int    `json:"id_imt" form:"id_imt"`
	ID_Kesehatan int    `json:"id_kesehatan" form:"id_kesehatan"`
	Nama         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
}

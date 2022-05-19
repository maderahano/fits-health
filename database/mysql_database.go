package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"fits-health/config"
	"fits-health/internal/entity"
)

func InitDB(conf config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err.Error())
	}

	initMigrate(DB)
	return DB
}

func initMigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.IMT{})
	db.AutoMigrate(&entity.ProfilKesehatan{})
	db.AutoMigrate(&entity.ResepMakanan{})
	db.AutoMigrate(&entity.JadwalMakanan{})
	db.AutoMigrate(&entity.Artikel{})
	db.AutoMigrate(&entity.Olahraga{})
}

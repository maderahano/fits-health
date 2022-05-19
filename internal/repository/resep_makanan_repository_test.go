package repository

import (
	"database/sql/driver"
	"fits-health/internal/entity"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateResep(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlResepRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
		WithArgs("Nasi Goreng", 30, 150, 10, "Misal Bahan", "Misal Persiapan", "Misal Langkah", 1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.CreateResep(entity.ResepMakanan{
		ID:        1,
		Nama:      "Nasi Goreng",
		Durasi:    30,
		Kalori:    150,
		Porsi:     10,
		Bahan:     "Misal Bahan",
		Persiapan: "Misal Persiapan",
		Langkah:   "Misal Langkah",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetAllResep(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlResepRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `resep_makanans`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nama", "durasi", "kalori", "porsi", "bahan", "persiapan", "langkah"}).
			AddRow(1, "Nasi Goreng", 30, 150, 1, "Misal Bahan 1", "Misal Persiapan 1", "Misal Langkah 1").
			AddRow(2, "Mie Goreng", 10, 125, 1, "Misal Bahan 2", "Misal Persiapan 2", "Misal Langkah 2").
			AddRow(3, "Ayam Goreng", 45, 175, 3, "Misal Bahan 3", "Misal Persiapan 3", "Misal Langkah 3"))

	res := repo.GetAllResep()
	assert.Equal(t, res[2].Nama, "Ayam Goreng")
	assert.Len(t, res, 3)
}

func TestGetResepByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlResepRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `resep_makanans`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nama", "durasi", "kalori", "porsi", "bahan", "persiapan", "langkah"}).
			AddRow(1, "Nasi Goreng", 30, 150, 1, "Misal Bahan", "Misal Persiapan", "Misal Langkah"))

	res, err := repo.GetResepByID(1)
	assert.Equal(t, res.Nama, "Nasi Goreng")
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateResepByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlResepRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("Mie Goreng", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateResepByID(1, entity.ResepMakanan{
		Nama: "Mie Goreng",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteResepByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlResepRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteResepByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

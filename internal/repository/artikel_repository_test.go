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

func TestCreateArtikel(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlArtikelRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
		WithArgs("Penyebab Tubuh Kamu Gendut", "Kesehatan", "misal gambar", "misal isi", 1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.CreateArtikel(entity.Artikel{
		ID:     1,
		Judul:  "Penyebab Tubuh Kamu Gendut",
		Jenis:  "Kesehatan",
		Gambar: "misal gambar",
		Isi:    "misal isi",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetAllArtikel(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlArtikelRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `artikels`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "judul", "jenis", "gambar", "isi"}).
			AddRow(1, "Penyebab Tubuh Kamu Gendut", "Kesehatan", "misal gambar 1", "misal isi 1").
			AddRow(2, "Penyebab Tubuh Kamu Kurus", "Kesehatan", "misal gambar 2", "misal isi 2"))

	res := repo.GetAllArtikel()
	assert.Equal(t, res[1].Judul, "Penyebab Tubuh Kamu Kurus")
	assert.Len(t, res, 2)
}

func TestGetArtikelByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlArtikelRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `artikels`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "judul", "jenis", "gambar", "isi"}).
			AddRow(1, "Penyebab Tubuh Kamu Gendut", "Kesehatan", "misal gambar", "misal isi"))

	res, err := repo.GetArtikelByID(1)
	assert.Equal(t, res.Judul, "Penyebab Tubuh Kamu Gendut")
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateArtikelByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlArtikelRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("Penyebab Tubuh Kamu Normal", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateArtikelByID(1, entity.Artikel{
		Judul: "Penyebab Tubuh Kamu Normal",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteArtikelByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlArtikelRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteArtikelByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

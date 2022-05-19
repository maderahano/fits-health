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

func TestCreateJadwal(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlJadwalRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
		WithArgs(1, "12 Desember 2022", "Siang", 1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.CreateJadwal(entity.JadwalMakanan{
		ID:       1,
		ID_Resep: 1,
		Tanggal:  "12 Desember 2022",
		Waktu:    "Siang",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetAllJadwal(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlJadwalRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `jadwal_makanans`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "id_resep", "tanggal", "waktu"}).
			AddRow(1, 1, "12 Desember 2022", "Siang").
			AddRow(2, 6, "12 Desember 2022", "Malam"))

	res := repo.GetAllJadwal()
	assert.Equal(t, res[1].Waktu, "Malam")
	assert.Len(t, res, 2)
}

func TestGetJadwalByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlJadwalRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `jadwal_makanans`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "id_resep", "tanggal", "waktu"}).
			AddRow(1, 3, "12 Desember 2022", "Siang"))

	res, err := repo.GetJadwalByID(1)
	assert.Equal(t, res.Waktu, "Siang")
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateJadwalByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlJadwalRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(2, "12 Desember 2022", "Pagi", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateJadwalByID(1, entity.JadwalMakanan{
		ID:       1,
		ID_Resep: 2,
		Tanggal:  "12 Desember 2022",
		Waktu:    "Pagi",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteJadwalByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlJadwalRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteJadwalByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

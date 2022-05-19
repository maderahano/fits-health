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

func TestCreateProfilKesehatan(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlProfilKesehatanRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
		WithArgs(80, 180, "AB", "160/80", 1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.CreateProfilKesehatan(entity.ProfilKesehatan{
		ID:            1,
		BeratBadan:    80,
		TinggiBadan:   180,
		GolonganDarah: "AB",
		Tensi:         "160/80",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetAllProfilKesehatan(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlProfilKesehatanRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `profil_kesehatans`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "berat_badan", "tinggi_badan", "golongan_darah", "tensi"}).
			AddRow(1, 80, 180, "AB", "160/80").
			AddRow(2, 56, 173, "B", "180/56").
			AddRow(3, 98, 177, "A", "155/60"))

	res := repo.GetAllProfilKesehatan()
	assert.Equal(t, res[1].TinggiBadan, 173)
	assert.Len(t, res, 3)
}

func TestGetProfilKesehatanByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlProfilKesehatanRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `profil_kesehatans`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "berat_badan", "tinggi_badan", "golongan_darah", "tensi"}).
			AddRow(1, 80, 180, "AB", "160/80"))

	res, err := repo.GetProfilKesehatanByID(1)
	assert.Equal(t, res.GolonganDarah, "AB")
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateProfilKesehatanByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlProfilKesehatanRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(83, 185, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateProfilKesehatanByID(1, entity.ProfilKesehatan{
		BeratBadan:  83,
		TinggiBadan: 185,
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteProfilKesehatanByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlProfilKesehatanRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteProfilKesehatanByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

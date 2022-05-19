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

func TestCreateOlahraga(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlOlahragaRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
		WithArgs("Push Up", "Misal Isi Deskripsi Push Up", 1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.CreateOlahraga(entity.Olahraga{
		ID:        1,
		Nama:      "Push Up",
		Deskripsi: "Misal Isi Deskripsi Push Up",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetAllOlahraga(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlOlahragaRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `olahragas`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nama", "deskripsi"}).
			AddRow(1, "Push Up", "Misal Isi Deskripsi Push Up").
			AddRow(2, "Sit Up", "Misal Isi Deskripsi Sit Up").
			AddRow(3, "Pull Up", "Misal Isi Deskripsi Pull Up"))

	res := repo.GetAllOlahraga()
	assert.Equal(t, res[2].Nama, "Pull Up")
	assert.Len(t, res, 3)
}

func TestGetOlahragaByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlOlahragaRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `olahragas`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nama", "deskripsi"}).
			AddRow(1, "Push Up", "Misal isi Deskripsi Push Up"))

	res, err := repo.GetOlahragaByID(1)
	assert.Equal(t, res.Nama, "Push Up")
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateOlahragaByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlOlahragaRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("Push Up", "Misal Isi Deskripsi Push Up", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOlahragaByID(1, entity.Olahraga{
		Nama:      "Push Up",
		Deskripsi: "Misal Isi Deskripsi Push Up",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteOlahragaByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlOlahragaRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteOlahragaByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

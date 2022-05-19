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

func TestCreateIMT(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlIMTRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
		WithArgs("Hyper Obesitas", 1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.CreateIMT(entity.IMT{
		ID:   1,
		Nama: "Hyper Obesitas",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetAllIMT(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlIMTRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `imts`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nama"}).
			AddRow(1, "Super Obesitas").
			AddRow(2, "Hyper Obesitas"))

	res := repo.GetAllIMT()
	assert.Equal(t, res[0].Nama, "Super Obesitas")
	assert.Len(t, res, 2)
}

func TestGetIMTByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlIMTRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `imts`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nama"}).
			AddRow(1, "Super Obesitas"))

	res, err := repo.GetIMTByID(1)
	assert.Equal(t, res.Nama, "Super Obesitas")
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateIMTByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlIMTRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("Hyper Obesitas", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateIMTByID(1, entity.IMT{
		Nama: "Hyper Obesitas",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteIMTByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlIMTRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteIMTByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

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

func TestRegisterUser(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
		WithArgs(2, 1, "rahano", "rahano@gmail.com", "123456", 1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.RegisterUser(entity.User{
		ID:           1,
		ID_IMT:       2,
		ID_Kesehatan: 1,
		Nama:         "rahano",
		Email:        "rahano@gmail.com",
		Password:     "123456",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetAllUsers(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "id_imt", "id_kesehatan", "nama", "email", "password"}).
			AddRow(1, 2, 1, "Rahano", "rahano@gmail.com", "123456").
			AddRow(2, 1, 2, "Fabyan", "fabyan@gmail.com", "123456").
			AddRow(3, 3, 3, "Zaid", "zaid@gmail.com", "123456"))

	res := repo.GetAllUsers()
	assert.Equal(t, res[0].Nama, "Rahano")
	assert.Len(t, res, 3)
}

func TestGetUserByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "id_imt", "id_kesehatan", "nama", "email", "password"}).
			AddRow(1, 2, 1, "Rahano", "rahano@gmail.com", "123456"))

	res, err := repo.GetUserByID(1)
	assert.Equal(t, res.Nama, "Rahano")
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestGetOneByEmail(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "id_imt", "id_kesehatan", "nama", "email", "password"}).
			AddRow(1, 2, 1, "Rahano", "rahano@gmail.com", "123456"))

	res, err := repo.GetUserByEmail("rahano@gmail.com")
	assert.Equal(t, res.Nama, "Rahano")
	assert.NoError(t, err)
	assert.True(t, true)
}
func TestUpdateOneByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("rahano", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateUserByID(1, entity.User{
		Nama: "rahano",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestDeleteOneByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	dial := mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}}
	db, _ := gorm.Open(dial)
	repo := NewMysqlUserRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteUserByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

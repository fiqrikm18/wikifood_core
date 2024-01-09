package showcase

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestNewShowcaseRepository(t *testing.T) {
	dbMock, _, err := sqlmock.New()
	assert.Nil(t, err)

	db, _ := gorm.Open(postgres.New(postgres.Config{
		Conn:       dbMock,
		DriverName: "postgres",
	}), &gorm.Config{})
	repository := NewShowcaseRepository(db)
	assert.NotNil(t, repository)

}

func TestShowcaseRepository_CreateShowcase(t *testing.T) {
	dbMock, _, err := sqlmock.New()
	assert.Nil(t, err)

	db, _ := gorm.Open(postgres.New(postgres.Config{
		Conn:       dbMock,
		DriverName: "postgres",
	}), &gorm.Config{})
	repository := NewShowcaseRepository(db)
	assert.NotNil(t, repository)
}

func TestShowcaseRepository_UpdateShowcase(t *testing.T) {

}

func TestShowcaseRepository_DeleteShowcase(t *testing.T) {

}

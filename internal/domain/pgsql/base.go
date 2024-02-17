package pgsql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Dsn string   `json:"dsn"`
	DB  *gorm.DB `json:"db"`
}

func InitDB(dsn string) (*DB, error) {
	var pg DB
	pg.Dsn = dsn

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &pg, err
	}
	pg.DB = db

	err = db.AutoMigrate(
	// TODO: Some models for create/update
	)
	if err != nil {
		return &pg, err
	}

	return &pg, nil
}

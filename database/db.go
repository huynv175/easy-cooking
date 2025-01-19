package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	mysqldb, err := d.db.DB()
	if err != nil {
		return err
	}
	if err := mysqldb.Close(); err != nil {
		return err
	}
	return nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}

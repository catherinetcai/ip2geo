package database

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	*gorm.DB
}

func New() (*DB, error) {
	cfg := &mysql.Config{
		User:   "root",
		Passwd: "",
		Addr:   "localhost:4000",
		DBName: "test",
	}
	db, err := gorm.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &DB{db}, err
}

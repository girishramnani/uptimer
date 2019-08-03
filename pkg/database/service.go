package database

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type DBManager struct {
	*pg.DB
}

func NewDBManager(options *pg.Options) *DBManager {
	return &DBManager{
		pg.Connect(options),
	}
}

// CreateSchema creates the table for the struct provided if it doesn't exist already
func (dbm *DBManager) CreateSchema(model interface{}) error {
	err := dbm.CreateTable(model, &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	return nil
}

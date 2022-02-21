// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package pgdb_client

import (
	"github.com/go-pg/pg/v10"
	"log"
)

type DBConfig struct {
	Prod             bool
	Addr             string
	User             string
	Password         string
	Database         string
	DBModel          []interface{}
	DBCompositeTypes []interface{}
}

type DBComponent struct {
	config *DBConfig
	db     *pg.DB
	prod   bool
}

func NewDBComponent(dbConfig *DBConfig) *DBComponent {
	mtd := "NewDBComponent: "
	// CIRA = Construction = Initialization = Return (Resource) Allocation

	// Nil check
	if dbConfig.DBModel == nil {
		msg := " NPE: DBSchema is NIl"
		dbgPrint(mtd + msg)
		log.Fatal(msg)
	}
	dbgPrint(mtd + " 1. Construction")
	dbComp := &DBComponent{
		config: dbConfig,
		prod:   dbConfig.Prod,
	}

	dbgPrint(mtd + "  2. Connect Database")
	dbComp.ConnectDataBase()

	if dbConfig.DBModel != nil { // if we don't have a schema, resume w/o creating a DB to keep init time short
		dbgPrint(mtd + " CreateDataBase")
		dbComp.CreateDataBase(dbConfig.DBCompositeTypes, dbConfig.DBModel)
	}

	dbgPrint(mtd + "  3. Return (Reference) to DBComp")

	return dbComp
}

func (c *DBComponent) ConnectDataBase() {
	db := pg.Connect(&pg.Options{
		Addr:     c.config.Addr,
		User:     c.config.User,
		Password: c.config.Password,
		Database: c.config.Database,
	})
	c.db = db
}

// PingDataBase returns true / false if the DB can be reached.
func (c *DBComponent) PingDataBase() (bool, error) {
	db := c.db
	err := db.Ping(db.Context())
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Shutdown closes the DB
func (c *DBComponent) Shutdown() error {
	err := c.db.Close()
	if err != nil {
		return err
	} else {
		return nil
	}
}

// DB returns the DB
func (c *DBComponent) DB() *pg.DB {
	return c.db
}

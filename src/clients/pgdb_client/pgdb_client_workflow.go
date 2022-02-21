package pgdb_client

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
)

// CreateDataBase creates a DB for the given DB config
func (c *DBComponent) CreateDataBase(compositeTypes []interface{}, schema []interface{}) *pg.DB {
	mtd := "CreateDataBase: "

	if compositeTypes != nil {
		dbgPrint(mtd + " CreateCompositeTypes")
		compErr := c.CreateCompositeTypes(compositeTypes)
		if compErr != nil {
			log.Println(mtd + "Can't create or update DB composite types")
			log.Fatal(compErr)
		}
	}

	if schema != nil {
		dbgPrint(mtd + " CreateSchema")
		dbCreateErr := c.CreateSchema(schema)
		if dbCreateErr != nil {
			log.Println(mtd + "Can't create or update DB schema")
			log.Fatal(dbCreateErr)
		}
	}
	return c.db
}

// CreateCompositeTypes creates composite types for the supplied array of structs.
// Must be called before creating the db schema to ensure all required types are present
func (c *DBComponent) CreateCompositeTypes(models []interface{}) error {
	mtd := "CreateCompositeTypes: "
	db := c.db

	dbgPrint(mtd + " Teardown happens in reverse order of creation.")
	for i := len(models) - 1; i >= 0; i-- {
		err := db.Model(models[i]).DropComposite(&orm.DropCompositeOptions{
			IfExists: true,
			Cascade:  true,
		})
		if err != nil {
			return err
		}
	}

	dbgPrint(mtd + " Type creation in actual order!")
	for _, model := range models {
		err := db.Model(model).CreateComposite(nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateSchema creates database schema for the supplied array of structs.
func (c *DBComponent) CreateSchema(models []interface{}) error {
	mtd := "CreateSchema: "
	db := c.db

	dbgPrint(mtd + " Schema creation in actual order!")
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        !c.prod, // If prod, make persistent tables, else make temp tables
			IfNotExists: c.prod,  // if prod, only create tables if not exists already
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateManyToManyTables creates M2M relation tables for the supplied array of structs.
// Must be called after CreateSchema to ensure the presence of all tables
func (c *DBComponent) CreateManyToManyTables(models []interface{}) error {

	db := c.db
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
			Temp:        true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

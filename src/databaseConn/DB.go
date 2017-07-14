package databaseConn

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)
type DB struct {
}
var (db *gorm.DB)

func init() {
	db = DB{}.GetDB()
}

func (database DB) GetDB() *gorm.DB{
	var err error
	db, err = gorm.Open("sqlite3", "test.sqlite")
	if err != nil{
		log.Fatal(err)
	}
	return db
}


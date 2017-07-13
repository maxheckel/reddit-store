package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"fmt"
)

var (db *gorm.DB)

func init() {
	db = getDB()
}

func getDB() *gorm.DB{
	var err error
	db, err = gorm.Open("sqlite3", "test.sqlite")
	if err != nil{
		fmt.Println("blah")
		log.Fatal(err)
	}



	return db
}


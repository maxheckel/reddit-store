package main

import (
	"models"
	"seeds"
	//"fmt"
	"databaseConn"
)

func main() {
	var db = databaseConn.DB{}.GetDB()
	defer db.Close()
	err := db.AutoMigrate(&models.Comment{}, &models.Post{}, &models.Image{}, &models.Subreddit{})
	if err.Error != nil{
		panic("Migrating failed")
	}
	seeds.Execute()
}

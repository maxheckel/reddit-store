package main

import (
	"models"
	"fmt"
)

func main() {
	defer db.Close()
	err := db.AutoMigrate(&models.Comment{}, &models.Post{}, &models.Image{})
	fmt.Printf("%+v\n", err)


}

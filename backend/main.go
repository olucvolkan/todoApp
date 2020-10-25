package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/olucvolkan/todoApp/config"
	"github.com/olucvolkan/todoApp/models"
	"github.com/olucvolkan/todoApp/routes"
)

var err error

func main() {

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Todo{})

	r := routes.SetupRouter()
	// running
	r.Run()
}

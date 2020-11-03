package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/olucvolkan/todoApp/config"
	"github.com/olucvolkan/todoApp/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/olucvolkan/todoApp/models"
	"github.com/olucvolkan/todoApp/routes"
	"github.com/subosito/gotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"os"
	"database/sql"
	
)

var err error

// @title Todo Api Example
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email volkanoluc@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @BasePath /
// @query.collection.format multi
// @x-extension-openapi {"example": "value on a json format"}

func main() {
	gotenv.Load()
	createDBDsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		3306,
	)

	db, err := sql.Open("mysql",createDBDsn)
	db.Exec("CREATE DATABASE IF NOT EXISTS " +os.Getenv("DB_NAME") + ";")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database or updated")
	}
	db.Close()
	
	config.DB, _ = gorm.Open("mysql", config.DbURL())
	config.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Todo{})
	defer config.DB.Close()
	
	r := routes.SetupRouter()
	// running


	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Todo Api Example"
	docs.SwaggerInfo.Description = "This is a sample Todo app server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.Schemes = []string{"http"}


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8081")
}

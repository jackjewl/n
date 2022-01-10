package db

import (
	"fmt"
	"github.com/jackjewl/squirrelBlog/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB gorm.DB

func InitDb() {

	config := utils.InitConfig()

	dsn := config.Mysql.Username + ":" + config.Mysql.Password + "@tcp(" + config.Mysql.Url + ")/" + config.Mysql.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}
	DB = *db
}

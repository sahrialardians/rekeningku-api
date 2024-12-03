package configs

import (
	"github.com/sahrialardians/rekeningku/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	dsn := "user:pass@tcp(127.0.0.1:3306)/rekeningku?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utils.ErrorPanic(err)

	return db
}

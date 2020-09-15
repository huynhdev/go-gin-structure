package models

import (
	"fmt"
	"log"

	"github.com/huynhdev/go-gin-structure/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	gorm.Model
}

// Setup initializes the database instance
func Setup() {
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	db.AutoMigrate(
		&User{},
		&Tag{},
		&Article{},
	)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}

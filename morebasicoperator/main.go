package main

import (
	"os"

	"morebasicoperator/model"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	user model.User
	log  = logrus.New()
	db   *gorm.DB
)

func init() {
	log.Out = os.Stdout
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	db.Create(&model.User{})

	log.Info(user)

	if err := db.Where("Number = ?", 16701).Find(&user).Error; err != nil {
		log.Panic(err)
	}

	log.Info(user.Name)
	log.Info(user.Number)
	log.Info(user)
}

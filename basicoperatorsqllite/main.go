package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name   string
	Number uint
}

var (
	log = logrus.New()
)

func init() {
	log.Out = os.Stdout
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create//.Find(&users)
	//db.Create(&Product{Name: "MirChen", Number: 16701})

	// Read
	var product Product
	result := db.Find(&product) // find product with code D42
	log.Info(result)

	log.Info(product.Name)
	log.Info(product.Number)
	log.Info(product)

	// Update - update product's price to 200
	db.Model(&product).Update("Number", 200)

	log.Info(product.Name)
	log.Info(product.Number)
	log.Info(product)

	// Update - update multiple fields
	db.Model(&product).Updates(Product{Number: 201, Name: "Test"}) // non-zero fields

	log.Info(product.Name)
	log.Info(product.Number)
	log.Info(product)

	db.Model(&product).Updates(map[string]interface{}{"Number": 202, "Name": "MirChenTest"})

	log.Info(product.Name)
	log.Info(product.Number)
	log.Info(product)
}

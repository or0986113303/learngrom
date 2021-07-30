package control

import (
	"log"
	"morebasicoperator/model"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ORMEngine struct {
	*gorm.DB
}

var (
	ormEngine *ORMEngine
	ormOnce   sync.Once
)

func NewEngine() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
	}

	ormOnce.Do(func() {
		dbh := new(ORMEngine)
		dbh.DB = db
		ormEngine = dbh
	})
}

func GetEngine() *ORMEngine {
	return ormEngine
}

func (USR *ORMEngine) Select(unmber int) (res model.User) {
	if err := USR.Where("Number = ?", unmber).Find(&res).Error; err != nil {
		log.Panic(err)
	}
	return
}

package control

import (
	"cobragingorm/internal/pkg/model"
	"reflect"
	"sync"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ORMEngineInterface interface {
	Select() model.User
}

type ORMEngine struct {
	*gorm.DB
}

var (
	ormEngine *ORMEngine
	ormOnce   sync.Once
	log       = logrus.New()
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

func (USR *ORMEngine) Insert(src model.User) (res model.User) {
	USR.AutoMigrate(&model.User{})
	USR.Create(&src)
	res = src
	return
}

func (USR *ORMEngine) SelectConditionLast(Name string, value interface{}) (res model.User) {

	cond := Name + " = ?"
	r := reflect.ValueOf(&res)
	switch reflect.Indirect(r).FieldByName(Name).Type().String() {
	case "uint":
		if err := USR.Where(cond, value.(int)).Last(&res).Error; err != nil {
			log.Panic(err)
		}
	case "string":
		if err := USR.Where(cond, value.(string)).Last(&res).Error; err != nil {
			log.Panic(err)
		}
	default:
		log.Panic("gg")
	}
	return
}

func (USR *ORMEngine) SelectConditionFirst(Name string, value interface{}) (res model.User) {

	cond := Name + " = ?"
	r := reflect.ValueOf(&res)
	switch reflect.Indirect(r).FieldByName(Name).Type().String() {
	case "uint":
		if err := USR.Where(cond, value.(int)).First(&res).Error; err != nil {
			log.Panic(err)
		}
	case "string":
		if err := USR.Where(cond, value.(string)).First(&res).Error; err != nil {
			log.Panic(err)
		}
	default:
		log.Panic("gg")
	}
	return
}

func (USR *ORMEngine) SelectCondition(Name string, value interface{}) (res model.User) {

	cond := Name + " = ?"
	r := reflect.ValueOf(&res)
	switch reflect.Indirect(r).FieldByName(Name).Type().String() {
	case "uint":
		if err := USR.Where(cond, value.(int)).Find(&res).Error; err != nil {
			log.Panic(err)
		}
	case "string":
		if err := USR.Where(cond, value.(string)).Find(&res).Error; err != nil {
			log.Panic(err)
		}
	default:
		log.Panic("gg")
	}
	return
}

func (USR *ORMEngine) OmitName(schema interface{}) (res model.User) {

	switch v := schema.(type) {
	case string:
		log.Panic(v)
	case int32, int64:
		log.Panic(v)
	case model.User:
		op, ok := schema.(model.User)
		if ok {
			USR.Omit("Name").Create(&op)
		} else {
			log.Panic("convert to other model fail")
		}
	default:
		log.Panic("unknown")
	}

	return
}

package main

import (
	"os"

	"morebasicoperator/control"
	"morebasicoperator/model"

	"github.com/sirupsen/logrus"
)

var (
	user model.User
	log  = logrus.New()
	dbh  *control.ORMEngine
)

func init() {
	log.Out = os.Stdout
	control.NewEngine()
}

func main() {
	dbh = control.GetEngine()

	src := model.User{Name: "Mir", Number: 16704}

	res := dbh.Insert(src)
	log.Info(res) // there should to get output which number is 16701 and name is Mir

	res = dbh.SelectCondition("Number", 16702)
	log.Info(res) // there should to get output which number is 16701

	dbh.OmitName(src)

	res = dbh.SelectConditionFirst("Name", "Mir")
	log.Info(res) // there should to get output which number is 16701

	res = dbh.SelectConditionLast("Name", "Mir")
	log.Info(res) // there should to get output which number is 16701

}

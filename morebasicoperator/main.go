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
	res := dbh.Select(16701)
	log.Info(res)
}

package main

import (
	"os"

	"morebasicoperator/model"

	"github.com/sirupsen/logrus"
)

var (
	user = model.User{Name: "Mir", Number: 16701}
	log  = logrus.New()
)

func init() {
	log.Out = os.Stdout
}

func main() {
	log.Info(user)
}

package control

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"

	"cobragingorm/internal/pkg/control"
	"cobragingorm/internal/pkg/model"
)

var (
	log = logrus.New()
)

// @Tags user
// @Summary put usr object
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body []model.User false "Name and Number of user"
// @Success 200 {nil} nil "Insert result success or not"
// @Router /auth/api/v1/usr/info [put]
func UsrInfo(context *gin.Context) {
	control.NewEngine()
	dbh := control.GetEngine()
	var data model.User
	context.BindJSON(&data)
	// src := model.User{Name: gjson.Get(value, "Name").String(), Number: uint(gjson.Get(value, "Number").Uint())}
	res := dbh.Insert(data)
	log.Info(res)
	context.JSON(http.StatusOK, nil)
}

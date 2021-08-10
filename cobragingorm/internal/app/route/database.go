package route

import (
	"cobragingorm/api/control"

	"github.com/gin-gonic/gin"
)

func InitDataBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	DatabaseRouter := Router.Group("database")
	{
		DatabaseRouter.PUT("usrinfo", control.UsrInfo)
	}
	return DatabaseRouter
}

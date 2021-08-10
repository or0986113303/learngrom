package route

import (
	"github.com/gin-gonic/gin"
)

func InitV1Router(Router *gin.RouterGroup) (R gin.IRoutes) {
	// v1 group route
	v1 := Router.Group("/api/v1/")
	{
		InitDataBaseRouter(v1)
	}

	return v1
}

package route

import (
	"github.com/gin-gonic/gin"
)

func InitV2Router(Router *gin.RouterGroup) (R gin.IRoutes) {
	// v1 group route
	v2 := Router.Group("/api/v2/")
	{
		InitDataBaseRouter(v2)
	}

	return v2
}

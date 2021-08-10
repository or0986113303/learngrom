package api

import (
	"cobragingorm/internal/app/route"
	"github.com/gin-gonic/gin"
)

// run http daemon
func HttpRun() {
	application := route.InitRouter()
	application.Use(gin.Logger())
	application.Use(gin.Recovery())
	application.Run("localhost:8080")
}

package route

import (
	"cobragingorm/api/middleware"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	log = logrus.New()
)

// InitRouter ...
func InitRouter() *gin.Engine {
	Router := gin.Default()

	// public group route
	PublicGroup := Router.Group("")
	{
		//InitWebRouter(PublicGroup)
		Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		InitV2Router(PublicGroup)
	}

	authMiddleware, err := middleware.InitJWTmiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	Router.POST("/login", authMiddleware.LoginHandler)

	auth := Router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		InitV1Router(auth)
	}
	return Router
}

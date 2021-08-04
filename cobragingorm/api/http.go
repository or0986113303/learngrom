package main

import (
	"cobragingorm/internal/app/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// run http daemon
func httpRun(conf *config.Config) {
	gin.SetMode(gin.ReleaseMode)
	// start daemon
	corsConf := conf.Cors
	router := api.NewDefault(
		cors.New(cors.Config{
			// AllowOrigins:     corsConf.Origin,
			AllowMethods:     corsConf.Methods,
			AllowHeaders:     corsConf.Headers,
			AllowCredentials: corsConf.Credentials,
			AllowOriginFunc: func(origin string) bool {
				for _, host := range corsConf.Origin {
					if origin == host || host == "*" {
						return true
					}
				}
				return false
			},
			MaxAge: 12 * time.Hour,
		}),
	)

	api.SetRoutes(router)

	log.Infof("API Server Listening: %s", conf.Server.Addr)
	if err := router.Run(conf.Server.Addr); err != nil {
		log.Fatalf("run api error: %v", err)
		panic(err)
	}
}

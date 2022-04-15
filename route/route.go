package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ippool_api/api"
	"ippool_api/middleware"
	. "ippool_api/utils/log"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders: []string{"*"},
	}))
	r.Use(gin.Recovery())
	r.Use(middleware.Log())
	r.GET("/peers", api.GetPeers)

	GlobalLog.Infof("Run server on %s", "0.0.0.0:60010")
	if err := r.Run("0.0.0.0:60010"); err != nil {
		GlobalLog.Fatalf("Run server %s", err.Error())
	}
}

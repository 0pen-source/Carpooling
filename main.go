package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/service"
	"github.com/gin-gonic/gin"
)

func main() {
	initializeConfiguration()
	dao.InitializeRedis(config.Redis.URL, config.Redis.PoolSize)
	dao.InitializeCache(config.CarpoolingDatabases)
	runtime.GOMAXPROCS(runtime.NumCPU())
	gin.SetMode(config.MODE)
	router := gin.Default()
	service.Checkcode = config.Checkcode
	router.Use(
		service.MarkTesting(),
	)
	developerGroup := router.Group("/v1/user")
	{
		developerGroup.POST("/phonetest", service.Phonetest)
		developerGroup.POST("/login", service.Login)
	}

	server := &http.Server{
		Addr:    config.Address,
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

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
	dao.InitializeConfiguration()
	dao.InitializeRedis()
	dao.InitializeCache()
	dao.InitOSSClient()
	runtime.GOMAXPROCS(runtime.NumCPU())
	gin.SetMode(dao.GetMODE())
	router := gin.Default()
	router.Use(
		service.MarkTesting(),
	)
	noLoginUserGroup := router.Group("/v1/user")
	{
		noLoginUserGroup.POST("/checkphone", service.Phonetest)
		noLoginUserGroup.POST("/login", service.Login)
		noLoginUserGroup.POST("/getcode", service.GetVerificationCode)
		noLoginUserGroup.POST("/checkcode", service.CheckCode)
		noLoginUserGroup.POST("/register", service.Register)
	}
	LoginUserGroup := router.Group("/v1/user")
	{
		LoginUserGroup.Use(
			service.Auth(),
		)
		LoginUserGroup.POST("/setinformation", service.SetInformation)
		LoginUserGroup.POST("/upload", service.Upload)
	}

	passengers := router.Group("/v1/passengers")
	{
		passengers.Use(
			service.Auth(),
		)
		passengers.POST("/creattrip", service.CreatTrip)
		passengers.POST("/upload", service.Upload)
	}

	server := &http.Server{
		Addr:    dao.GetAddress(),
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

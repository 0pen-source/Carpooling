package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/service/common"
	"github.com/0pen-source/Carpooling/service/drivers"
	"github.com/0pen-source/Carpooling/service/passengers"
	"github.com/0pen-source/Carpooling/service/user"
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
		common.MarkTesting(),
	)
	noLoginUserGroup := router.Group("/v1/user")
	{
		noLoginUserGroup.POST("/checkphone", user.Phonetest)
		noLoginUserGroup.POST("/login", user.Login)
		noLoginUserGroup.POST("/getcode", user.GetVerificationCode)
		noLoginUserGroup.POST("/checkcode", user.CheckCode)
		noLoginUserGroup.POST("/register", user.Register)
	}
	LoginUserGroup := router.Group("/v1/user")
	{
		LoginUserGroup.Use(
			common.Auth(),
		)
		LoginUserGroup.POST("/setinformation", user.SetInformation)
		LoginUserGroup.POST("/upload", user.Upload)
	}

	passenger := router.Group("/v1/passengers")
	{
		passenger.Use(
			common.Auth(),
		)
		passenger.POST("/creattrip", passengers.CreatTrip)
		passenger.POST("/upload", user.Upload)
		passenger.POST("/index", passengers.Index)
		passenger.POST("/searchtrip", passengers.SearchTrip)

	}

	driver := router.Group("/v1/driver")
	{
		driver.Use(
			common.Auth(),
		)
		driver.POST("/creattrip", drivers.CreatTrip)
		driver.POST("/upload", user.Upload)
		driver.POST("/index", drivers.Index)
		driver.POST("/searchtrip", drivers.SearchTrip)
	}

	server := &http.Server{
		Addr:    dao.GetAddress(),
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
